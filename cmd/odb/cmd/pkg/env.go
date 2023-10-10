package pkg

func Env() string {
	return `
# container
IMAGE="mysql:latest"
CONTAINER="autodb"

# migration
SQL_FILE="sample.sql"

# database
# dsn format: root:12345@tcp(localhost:3342)/autodb?charset=utf8
DB_DRIVER="mysql"
DB_USER="root"
DB_PASSWORD="12345"
DB_HOST="localhost"
DB_PORT="3343"
DB_NAME="autodb"


# TODO: generate types
`
}
