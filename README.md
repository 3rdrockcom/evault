# eVault
eVault is system that stores and retrieves a value.

## Instructions
```bash
# Perform Database Migrations
./goose -dir=./migrations mysql "username:password@tcp(localhost:3306)/evault?parseTime=true" up;

# Build application
go build -o evault;

# Run application
./evault -dsn="username:password@tcp(localhost:3306)/evault?parseTime=true";

# Create a user
./evault -dsn="username:password@tcp(localhost:3306)/evault?parseTime=true" -createUser -username="EPOINT" -password="iJvIwBs1uh" -programID=1;
```

## API Documentation
API documentation is available on Postman (Collection: eVault).