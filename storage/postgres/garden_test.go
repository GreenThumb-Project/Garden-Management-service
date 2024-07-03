package postgres

import (
	"fmt"
	pb "garden-managment-service/generated/gardenManagement"
	"reflect"
	"testing"
)

func TestCreateGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	garden := NewGardenManagementRepo(db)

	respcreate, err := garden.CreateGarden(&pb.CreateGardenRequest{
		Id:      "92a97aa1-a9d3-4b30-b24e-3681995ee86d",
		UserId:  "00e6f3c0-d302-4877-8e35-984ab1ae3c81",
		Name:    "ANY",
		AreaSqm: 12.2})

	fmt.Println(err)
	waitcreate := pb.CreateGardenResponse{
		Success: true,
	}

	if !reflect.DeepEqual(respcreate, &waitcreate) {
		t.Errorf("have %v , wont %v", respcreate, &waitcreate)
	}
}
