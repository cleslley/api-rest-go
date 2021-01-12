FROM golang:latest

WORKDIR /app

EXPOSE 3000

COPY go.mod .

COPY go.sum .

RUN go mod download 

COPY . .

RUN  go build 

CMD ["./api-rest-go"]