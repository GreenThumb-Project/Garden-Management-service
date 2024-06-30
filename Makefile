CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:03212164@localhost:5432/garden_management_service?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path databases/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path databases/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path databases/migrations -database '${DBURL}' -verbose force 1

mig-create-type:
	migrate create -ext sql -dir databases/migrations -seq create_types_enum

mig-create-gardens:
	migrate create -ext sql -dir databases/migrations -seq create_gardens_table

mig-create-plants:
	migrate create -ext sql -dir databases/migrations -seq create_plants_table 

mig-create-carelogs:
	migrate create -ext sql -dir databases/migrations -seq create_care_logs_table