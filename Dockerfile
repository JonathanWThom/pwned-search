FROM golang:alpine
WORKDIR /pwned
COPY . .
RUN go build -o pwned
ENTRYPOINT ["./pwned"]
