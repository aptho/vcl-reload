FROM golang:1.15.7 as builder
WORKDIR /root
COPY . .
RUN go build -o reload main.go

FROM varnish:latest
WORKDIR /root/
COPY --from=builder /root .
CMD ["./scripts/start.sh"]
