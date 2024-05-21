FROM golang:latest AS builder

WORKDIR /app
COPY ./shared ./shared
COPY ./services/auth ./services/auth

WORKDIR ./services/auth
RUN go mod download

WORKDIR ./cmd
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR ./services/auth/cmd
CMD ["./main"]
