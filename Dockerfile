FROM golang:1.17

WORKDIR /usr/src/app

COPY . .

CMD go run app.go