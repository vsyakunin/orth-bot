FROM golang:alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /build
COPY . .
RUN go build -o main .

FROM scratch

COPY --from=builder etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/main /
COPY --from=builder /build/prayers.json /
COPY --from=builder /build/icons /icons
COPY --from=builder /build/stats /stats

ENTRYPOINT ["./main"]