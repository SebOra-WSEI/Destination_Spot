FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
CMD ["./main"]
