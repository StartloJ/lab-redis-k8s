FROM golang:1.16-alpine3.14 

WORKDIR /app

COPY app/ .

CMD ["go", "run", "main.go"]