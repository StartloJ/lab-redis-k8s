FROM golang:1.16-alpine3.14 

WORKDIR /app

COPY app/ .

RUN go mod tidy && \
    go build -o app

CMD ["./app"]