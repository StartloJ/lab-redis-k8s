FROM golang:1.16-alpine3.14 AS build-env

WORKDIR /app

COPY app/ .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -i -a -o app .

FROM gcr.io/distroless/base
COPY --from=build-env /app/app /
CMD ["/app"]