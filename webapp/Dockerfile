# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.19.3-alpine3.16 as build

WORKDIR /app

# Download necessary Go modules
COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go build -o /webapp

## Deploy
FROM alpine:3.16.2
# Create a user
RUN adduser -D capricorn

# Tell docker that all future commands should run as the capricorn user
USER capricorn

COPY --from=build /webapp /webapp

EXPOSE 8080/tcp

CMD ["/webapp","8080"]