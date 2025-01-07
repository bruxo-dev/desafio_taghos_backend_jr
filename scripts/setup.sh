echo "Setting up environment variables..."
export DB_USER="postgres"
export DB_PASSWORD="123123"
export DB_NAME="desafio_taghos"
export DB_SSLMODE="disable"

echo "Installing dependencies..."
go mod tidy

echo "Initializing database..."
go run db/database.go

echo "Starting the application..."
go run main.go

echo "Setup complete."