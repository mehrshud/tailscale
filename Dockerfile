# Build stage for tailscale
FROM golang:alpine as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN apt-get update && apt-get install -y wireguard-tools

# Final stage for tailscale
FROM alpine:latest
RUN apk add --no-cache wireguard-tools
WORKDIR /root/
COPY --from=build /app/main .
CMD ["./main"]