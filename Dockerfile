# Use an official Go runtime as a parent image
FROM golang:1.16

# Set the working directory in the container
WORKDIR /app

# Copy the Go project source code into the container
COPY . .

# Install any dependencies (if needed)
RUN go get ./...

# Build the Go application
RUN go build -o main .

# Expose the port your Go application uses
EXPOSE 8080

# Run the start_server.sh script
CMD ["./start_server.sh"]
