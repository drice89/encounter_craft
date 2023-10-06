# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Install Gin using go get
# RUN go install github.com/gin-gonic/gin@latest

# Copy your Go application code to the container
COPY . .

# Build and run your Go application
CMD ["go", "run", "main.go"]
