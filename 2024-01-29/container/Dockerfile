# Inherit from parent container definition
FROM python:3.12-slim
# Run all future commands in /app directory
WORKDIR /app
RUN echo "Hello World!" > index.html
# Expose port 8080
EXPOSE 8080
# Run a simple HTTP server
CMD ["python3", "-u", "-m", "http.server", "8080"]
