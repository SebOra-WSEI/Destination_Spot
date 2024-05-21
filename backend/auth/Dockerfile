FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/cmd
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR /app/cmd
CMD ["./main"]
