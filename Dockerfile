FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY .. /var/www

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -installsuffix cgo -o app .

EXPOSE 8080
ENTRYPOINT ["/var/www/app"]