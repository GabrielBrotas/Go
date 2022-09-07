```bash
# command line vars
go run cmd/http/*.go -help

# install module
go get -u github.com/julienschmidt/httprouter

# run postgres db
docker run --name postgres -d -p 5432:5432 \
    -e POSTGRES_DB=go_movies \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres123 \
    postgres

# Run server
go run cmd/http/*.go
```