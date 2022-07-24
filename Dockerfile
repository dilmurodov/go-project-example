# install base image
FROM golang:1.18-alpine

# install dependencies

ENV $GOPATH=/go
ENV $PATH=$GOPATH/bin:$PATH

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./ ./

RUN go build -o /app/cmd cmd/main.go

# Default command
CMD [ "/app/cmd/main" ]