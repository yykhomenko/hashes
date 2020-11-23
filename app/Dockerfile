FROM golang:alpine
WORKDIR /go/src/github.com/cbi-sh/hashes/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 ./go/src/github.com/cbi-sh/hashes/main .
EXPOSE 5000
CMD ["./main"]
