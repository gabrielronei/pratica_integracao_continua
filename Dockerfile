FROM golang:alpine

WORKDIR /app

COPY . .

ENTRYPOINT ["go", "run", "math.go"]