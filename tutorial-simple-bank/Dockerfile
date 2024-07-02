FROM golang:1.21.11 AS builder
WORKDIR /app

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy the Go modules manifests
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main main.go

# Create a new stage from scratch
FROM golang:1.21.11

# Set the working directory
WORKDIR /app

# Install netcat-traditional and make for wait-for-it.sh and make commands
RUN apt-get update && apt-get install -y netcat-traditional make

# Ensure golang-migrate is installed and available in PATH
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy the compiled binary, migration files, and wait-for-it.sh script from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/db/migration ./db/migration
COPY Makefile .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

# Ensure the wait-for-it.sh and start.sh scripts are executable
RUN chmod +x /app/wait-for.sh
RUN chmod +x /app/start.sh

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/app/start.sh"]
