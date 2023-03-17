FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build

FROM alpine:latest

LABEL org.opencontainers.image.authors="Temuri Takalandze <me@abgeo.dev>"

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/assets ./assets
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/domains /usr/local/bin/domains

ENTRYPOINT ["domains"]
