FROM golang:1.20 as builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o ./build/dynu-updater ./cmd


FROM scratch

COPY --from=builder /src/build/dynu-updater /dynu-updater
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/dynu-updater"]
