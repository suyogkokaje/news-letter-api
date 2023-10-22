postgresinit:
	sudo docker run --name newsletter -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
postgres:
	sudo docker exec -it newsletter psql
createdb:
	sudo docker exec -it newsletter createdb --username=root --owner=root newsletter
dropdb:
	sudo docker exec -it newsletter dropdb newsletter

.PHONY: postgresinit postgres createdb dropdb