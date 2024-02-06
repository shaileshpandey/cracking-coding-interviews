# Use an official Golang runtime as the base image
FROM golang:1.17-alpine as builder

# Set the working directory in the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Use a lightweight base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the executable from the builder stage into the final image
COPY --from=builder /app/myapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]
