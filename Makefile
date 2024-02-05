.PHONY: migrate
migrate:
	psql -U postgres -d postgres -h localhost -a -f ./migrations/create_table.sql
	psql -U postgres -d postgres -h localhost -a -f ./migrations/insert_data.sql

serverStart:
	go run ../YellowPages/cmd/server/main.go

#clientStart:
#	go run ../YellowPages/cmd/client/main.go

