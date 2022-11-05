FROM golang:1.19.3
WORKDIR /app
COPY . .
RUN apt-get update && apt-get install libvips-dev -y && apt-get clean
RUN go build -o bin/server src/cmd/main.go
CMD ["./bin/server"]