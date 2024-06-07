# Build stage
FROM golang:1.22.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o buddytracker .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/cmd/server/buddytracker .

CMD ["./buddytracker"]
