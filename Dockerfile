FROM golang:latest
WORKDIR /go/src/github.com/cbi-sh/hashes/
#RUN go get -d -v golang.org/x/net/html
COPY ./main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/cbi-sh/hashes/main .
CMD ["./main"]
