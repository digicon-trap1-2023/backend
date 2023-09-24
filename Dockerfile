FROM golang:1.21.1-alpine3.17

WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]