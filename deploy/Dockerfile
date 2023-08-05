# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o crypto-charts .

# Run the tests in the container
FROM builder AS tester
RUN go test -v ./... --short

FROM alpine:latest

COPY --from=builder app/crypto-charts .
COPY --from=builder app/.crypto-charts.yaml .
COPY --from=builder app/pkg/uploader/gcp-credentials.json .
COPY --from=builder app/deploy/crontab .
COPY --from=builder app/deploy/run_job.sh .

RUN /usr/bin/crontab /crontab

# start cron
ENTRYPOINT /usr/sbin/crond -f -l 8
