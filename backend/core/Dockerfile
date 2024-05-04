FROM golang:latest AS builder
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod download
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR /app/cmd
CMD ["./main"]