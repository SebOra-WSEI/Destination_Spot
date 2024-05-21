FROM golang:latest AS builder
WORKDIR /app

# Copy shared package and core service
COPY ./shared ./shared
COPY ./services/core ./services/core

# Change working and download required external packages
WORKDIR ./services/core
RUN go mod download

# Change working and build main function
WORKDIR ./cmd
RUN go build -o main

FROM golang:latest
WORKDIR /app
COPY --from=builder /app ./
WORKDIR ./services/core/cmd
CMD ["./main"]
