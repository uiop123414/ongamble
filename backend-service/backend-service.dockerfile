# Use a base Ubuntu image
FROM ubuntu:20.04

# Set environment variables for version, OS, and architecture
ENV version=v4.14.1
ENV os=linux
ENV arch=amd64

# Install required dependencies (curl, tar, etc.)
RUN apt-get update && apt-get install -y curl tar

# Create /app directory in the container
RUN mkdir /app

# Copy the Go binary into the container
COPY backendApp /app

# Set the command to run the backend application
CMD ["/app/backendApp"]
