docker run -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 --name gator gator:latest
go install github.com/pressly/goose/v3/cmd/goose@latest

postgres://myuser:postgres@localhost:5432/gator
 goose postgres postgres://myuser:postgres@localhost:5432/gator up

 psql -U myuser -d gator

 docker run -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 --name gator gator-db:1.0
 