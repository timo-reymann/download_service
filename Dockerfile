FROM golang:alpine3.10 as builder
ADD src/ /build/
WORKDIR /build
RUN go build -o server .

FROM alpine:3.10
EXPOSE 8086
WORKDIR /app
COPY --from=builder /build/server server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/form.html form.html
CMD ["./server"]
