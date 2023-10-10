package pkg

func Env() string {
	return `
# container
IMAGE="mysql:latest"
CONTAINER="autodb"

# migration
SQL_FILE="sample.sql"

# database
# mysql dsn: 			root:12345@tcp(localhost:3342)/autodb?charset=utf8
# postgres dsn: 	postgres://postgres:12345@localhost:5432/autodb?sslmode=disable
DB_DRIVER="mysql"
DB_USER="root"
DB_PASSWORD="12345"
DB_HOST="localhost"
DB_PORT="3343"
DB_NAME="autodb"
DB_TABLES="8"

# generate types
TYPE_PKG_NAME="types"
`
}
