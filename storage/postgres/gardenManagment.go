package postgres

import (
	"database/sql"
	"fmt"
	pb "garden-managment-service/generated"
	pkg "garden-managment-service/package"
)

type GardenManagementRepo struct {
	DB *sql.DB
}

func NewGardenManagementRepo(db *sql.DB) *GardenManagementRepo {
	return &GardenManagementRepo{DB: db}
}

func (g *GardenManagementRepo) CreateGarden(in *pb.CreateGardenRequest) (*pb.CreateGardenResponces, error) {
	rows, err := g.DB.Exec(`
		INSERT 
		INTO 
		gardens(
			id,
			userId,
			name,
			area_sqm	
		)
		VALUES(
			$1,
			$2,
			$3,
			$4)
		`, in.Id, in.UserId, in.Name, in.AreaSqm)

	if err != nil {
		return &pb.CreateGardenResponces{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.CreateGardenResponces{Success: false}, err
	}

	return &pb.CreateGardenResponces{Success: true}, nil

}

func (g *GardenManagementRepo) ViewGarden(in *pb.ViewGardenRequest) (*pb.ViewGardenResponces, error) {
	var garden pb.ViewGardenResponces
	err := g.DB.QueryRow(`
			SELECT
				id,
				userId,
				name,
				area_sqm
			FROM gardens
			WHERE
				id=$1
				AND deleted_at=0
			`, in.Id).Scan(&garden.Id, &garden.UserId, &garden.Name, &garden.AreaSqm)

	return &garden, err

}

func (g *GardenManagementRepo) UpdateGarden(in *pb.UpdateGardenRequest) (*pb.UpdateGardenResponces, error) {

	params := make(map[string]interface{})
	var query = "UPDATE gardens SET "
	if in.UserId != "" {
		query += "user_id = :user_id, "
		params["user_id"] = in.UserId
	}
	if in.Name != "" {
		query += "name = :name, "
		params["name"] = in.Name
	}
	if in.Type != "" {
		query += "type = :type, "
		params["type"] = in.Type
	}
	if in.AreaSqm != 0 {
		query += "area_sqm = :area_sqm, "
		params["area_sqm"] = in.AreaSqm
	}

	query += "updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = in.Id
	query, args := pkg.ReplaceQueryParams(query, params)

	res, err := g.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return &pb.UpdateGardenResponces{Success: false}, fmt.Errorf("no rows affected, user with id %s not found", in.Id)
	}

	return &pb.UpdateGardenResponces{Success: true}, nil

}

func (g *GardenManagementRepo) DeleteGarden(in *pb.DeleteGardenRequest) (*pb.DeleteGardenResponces, error) {

	rows, err := g.DB.Exec(`
			UPDATE
				gardens
			SET de
				deleted_ad=date_part('epoch', current_timestamp)::INT 
			WHERE
				id=$1
		`, in.Id)

	if err != nil {
		return &pb.DeleteGardenResponces{Success: false}, err
	}

	rowsAffected, err := rows.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return &pb.DeleteGardenResponces{Success: false}, err
	}

	return &pb.DeleteGardenResponces{Success: true}, nil
}

func (g *GardenManagementRepo) ViewUserGardens(in *pb.ViewUserGardensRequest) (*pb.ViewUserGardensResponces, error) {

	rows, err := g.DB.Query(`
			select 
				id,
				user_id,
				name,
				type,
				area_sqm
			FROM
				gardens
			WHERE
				user_id=$1
			`, in.UserId)

	if err != nil {
		return nil, err
	}

	var gardens []*pb.Garden
	var garden pb.Garden

	for rows.Next() {

		err = rows.Scan(&garden.Id, &garden.UserId, &garden.Name, &garden.Type, &garden.AreaSqm)
		if err != nil {
			return nil, err
		}
		gardens = append(gardens, &garden)
	}
	return &pb.ViewUserGardensResponces{Gardens: gardens}, nil
}
