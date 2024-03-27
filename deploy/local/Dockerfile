FROM golang:1.21.8-alpine3.18 AS BuildStage

WORKDIR /app

# Initialize Go modules
COPY go.mod go.sum ./
RUN go mod download

# Install swag
RUN apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the rest of the application code
COPY . .

# Make the generate.sh script executable
RUN chmod +x bin/generate.sh

# Generate Swagger documentation
RUN bin/generate.sh

# Build the application
RUN CGO_ENABLED=0 go build -o mini-twitter

# Use a minimal base image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the BuildStage
COPY --from=BuildStage /app/mini-twitter /app/mini-twitter

# Set the entry point for the container
ENTRYPOINT ["./mini-twitter"]
