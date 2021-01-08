FROM golang:1.8-onbuild

WORKDIR /app

COPY . .

RUN go build -o app

ENTRYPOINT ["./app"]