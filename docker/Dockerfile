# Dockerfile

# Use the official Go image as the base
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and source code into the container
COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Build the Go application
RUN go build -o trainme .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application
CMD ["./trainme"]

