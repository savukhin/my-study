FROM golang:1.18.2-alpine as builder

WORKDIR /app-client

COPY . .
RUN go mod download
RUN go build -o /main .

CMD ["/main"]
