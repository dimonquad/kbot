FROM golang AS builder

WORKDIR /go/src/app
COPY . .    
RUN make build || true

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
ENTRYPOINT [ "./kbot" ]