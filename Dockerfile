## Compile
FROM golang:latest AS builder

WORKDIR /go/src/app

COPY . .

RUN make
#RUN --mount=type=cache,target=/root/.cache/go-build GOOS=linux GOARCH=amd64 go build -v -o bin/vc_issuer -ldflags "-w -s --extldflags '-static'" ./cmd/issuer/main.go

## Deploy
FROM debian:bookworm-slim

WORKDIR /

RUN apt-get update && apt-get install -y curl procps
RUN rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

CMD [ "./eduid_ladok" ]

