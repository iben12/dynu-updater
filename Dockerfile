FROM golang:1.20 as builder

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 go build -o /src/build/dynu-updater /src/cmd/...


FROM scratch

COPY --from=builder /src/build/dynu-updater /dynu-updater
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/dynu-updater"]
