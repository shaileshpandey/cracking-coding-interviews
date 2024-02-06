docker build -t myapp .
docker tag myapp myapp:v1.0
docker tag 1234567890ab myapp:v1.0
docker run -p 8080:8080 myapp