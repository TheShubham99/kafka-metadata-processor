# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
