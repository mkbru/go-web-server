FROM golang:1.12.3 as builder
WORKDIR /home/go-web-server
RUN git clone https://github.com/Fullscript/go-web-server.git .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-web-server .

FROM alpine:3.9
WORKDIR /app
COPY --from=builder /home/go-web-server .

RUN addgroup -S -g 123456 notrootgroup &&\
    adduser -S -u 123456 -G notrootgroup notrootuser

USER 123456:123456

EXPOSE 8080

ENTRYPOINT ["./go-web-server"]
