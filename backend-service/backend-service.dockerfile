# Use a base Ubuntu image
FROM alpine:latest

# Create /app directory in the container
RUN mkdir /app

# Copy the Go binary into the container
COPY backendApp /app

# Set the command to run the backend application
CMD ["/app/backendApp"]
