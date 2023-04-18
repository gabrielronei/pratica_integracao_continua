FROM golang:alpine

WORKDIR /app

COPY . .

ENTRYPOINT ["go", "test", "math.go", "math_test.go"]