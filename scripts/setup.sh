echo "Installing dependencies..."
go mod tidy

echo "Starting the application..."
go run cmd/main.go

echo "Setup complete."