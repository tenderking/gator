# Gator 
The Gator CLI is a command-line tool designed to manage and interact with RSS feeds. It allows users to register, add feeds, follow and unfollow feeds, and aggregate feed data. The tool is built with Go and uses PostgreSQL as its database backend. It is ideal for developers looking to integrate RSS feed management into their applications or workflows. This one of the projects on boot.dev to learn backend.

## Prerequisites

Make sure you have the following installed on your system:
- [PostgreSQL](https://www.postgresql.org/download/)
- [Go](https://golang.org/dl/)

## Installation

To install the gator CLI, run the following command:

```sh
go install
```

## Configuration

Set up your configuration file as needed. The `.gatorconfig.json` file should be created in your home folder with the PostgreSQL connection string and an empty `current_user_name`. Here's an example:

```json
{
  "db_url": "postgres://myuser:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

## Running the Program

```sh
gator reset
gator register <name>
gator addfeed "Hacker News RSS" "https://hnrss.org/newest"
gator following
gator unfollow "https://hnrss.org/newest"
gator login <name>
gator agg
gator feeds
gator users
gator browse 10
```

You can also use `go run .` instead of `gator` in development.

## Useful Commands in Development

To run the program, use the following commands:

```sh
goose postgres postgres://myuser:postgres@localhost:5432/gator up
psql -U myuser -d gator
```

## Using PostgreSQL with Docker
### Building Docker Image

To build the Docker image for the gator CLI, use the following command:

```sh
docker build -t gator:latest .
```

You can also run PostgreSQL using Docker. Here are some examples:

```sh
docker run -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 --name gator gator:latest
docker run -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 --name gator gator-db:1.0
```



## Commands

Here are a few commands you can run with the gator CLI in development:

- `goose up`: Apply all available migrations.
- `goose down`: Roll back the most recently applied migration.
- `goose status`: Print the status of all migrations.

```sh
goose postgres postgres://myuser:postgres@localhost:5432/gator up
```

For more information, refer to the [goose documentation](https://github.com/pressly/goose).

To access and run SQL commands in PostgreSQL:

```sh
psql -U myuser -d gator
```

