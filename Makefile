backend:
	go run main.go

db: 
	docker run --name aqua-farm-db -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres

init-db:
	cat ./initdb/query.sql | docker exec -i aqua-farm-db psql -U postgres -d postgres