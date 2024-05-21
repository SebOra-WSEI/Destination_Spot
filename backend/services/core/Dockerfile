FROM golang:latest AS builder

WORKDIR /app
COPY ./shared ./shared
COPY ./services/core ./services/core

WORKDIR ./services/core
RUN go mod download

WORKDIR ./cmd
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR ./services/core/cmd
CMD ["./main"]
