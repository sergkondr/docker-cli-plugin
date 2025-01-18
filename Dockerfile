FROM golang:1.22.1 as builder
ARG VERSION=dev
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-X 'main.version=$(VERSION)'" -o docker-cli-plugin /app/cmd/

FROM alpine:3.20
COPY --from=builder /app/docker-cli-plugin /app/docker-cli-plugin
CMD ["docker-cli-plugin"]
