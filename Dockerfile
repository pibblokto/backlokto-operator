FROM alpine

WORKDIR /app

COPY controller /app/controller

CMD ["/app/controller"]
