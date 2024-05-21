FROM golang:latest AS builder
WORKDIR /app

# Copy shared package and authorization service
COPY ./shared ./shared
COPY ./services/auth ./services/auth

# Change working and download required external packages
WORKDIR ./services/auth
RUN go mod download

# Change working and build main function
WORKDIR ./cmd
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR ./services/auth/cmd
CMD ["./main"]
