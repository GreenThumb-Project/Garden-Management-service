package postgres

import (
	"fmt"
	pb "garden-managment-service/generated/gardenManagement"
	"reflect"
	"testing"
)

func TestAddPlantGarden(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	plant := NewGardenPlantManagementRepo(db)

	resadd, err := plant.AddPlanttoGarden(&pb.AddPlanttoGardenRequest{
		Id:           "41278da8-9505-41f8-a338-be29e7db817b",
		GardenId:     "3c02ad1a-d7f6-438c-82fc-544b7fb1cf1e",
		Species:      "any",
		Quantity:     3,
		PlantingDate: "2024-01-01",
		Status:       "planned",
	})
	if err != nil {
		fmt.Println(err)
	}
	waitadd := pb.AddPlanttoGardenResponse{
		Success: true,
	}
	if !reflect.DeepEqual(resadd, &waitadd) {
		t.Errorf("have %v , wont %v", resadd, &waitadd)
	}

}
