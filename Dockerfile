FROM  golang:latest

WORKDIR /app/server

LABEL Authors="mrbelka12000"

LABEL version="1.0"

COPY go.mod .

COPY go.sum .

RUN go mod  download

COPY . .

RUN go build cmd/main.go

CMD ["./main"]