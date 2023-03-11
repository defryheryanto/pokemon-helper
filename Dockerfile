FROM golang:1.19-alpine

WORKDIR /

RUN mkdir /logs

WORKDIR /app

COPY . .
copy .env .
RUN go mod download

RUN go build -o /pokemon-helper ./cmd/...

EXPOSE 8080

CMD [ "/pokemon-helper" ]