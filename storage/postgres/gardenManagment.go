package postgres

import "database/sql"

type GardenManagementRepo struct {
	DB *sql.DB
}

func NewGardenManagementRepo(db *sql.DB) *GardenManagementRepo {
	return &GardenManagementRepo{DB: db}
}

func (g *GardenManagementRepo) CreateGarden() (error) {

	return nil
}

func (g *GardenManagementRepo) GetGardenByID() (error) {

	return nil
}

func (g *GardenManagementRepo) UpdateGarden() (error) {

	return nil
}

func (g *GardenManagementRepo) DeleteGarden() (error) {

	return nil
}