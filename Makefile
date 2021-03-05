postgres:
	docker run --name postgres13 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:13-alpine

psql:
	docker exec -it postgres13 psql -U root

.PHONY: postgres psql