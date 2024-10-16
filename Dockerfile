FROM golang:1.23.1-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "./main.go"]