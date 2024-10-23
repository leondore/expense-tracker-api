FROM golang:1.23-alpine

WORKDIR /expense-tracker-api
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/expense-tracker-api/bin/api"]
EXPOSE 8080
