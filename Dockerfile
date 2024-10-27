FROM golang:1.23-alpine

WORKDIR /expense-tracker-api
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api \
  && go build -o ./bin/migrate ./cmd/migrate

CMD ["/expense-tracker-api/bin/api"]
EXPOSE 8080
