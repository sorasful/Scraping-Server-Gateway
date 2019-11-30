FROM golang:latest

WORKDIR /app
COPY . .


RUN go build server.go


CMD ["./server"]