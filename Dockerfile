FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o rest-api ./cmd/main.go

CMD ["./rest-api"]