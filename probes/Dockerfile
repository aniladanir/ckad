FROM golang:1.25 as builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./

RUN go build -o ./bin/probes .

FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get install -y curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/probes /opt/app/probes

EXPOSE 3000

ENTRYPOINT ["/opt/app/probes"]