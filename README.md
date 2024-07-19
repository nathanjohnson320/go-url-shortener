### Setup

1. Install go-migrate https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
1. Run docker compose to create the db `docker compose up`
1. `migrate -source database/migrations -database postgres://localhost:5432/url_shortener up`

### Running

1. `go run main.go`
1. Navigate to `http://localhost:8080`
