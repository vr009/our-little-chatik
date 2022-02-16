FROM golang:latest

RUN go version

ENV GO111MODULE=on

ENV GOPATH=/

COPY ./ ./

EXPOSE 8080

RUN go mod download
RUN go mod tidy
RUN go build -o auth-service ./cmd/main.go


CMD ["./auth-service"]