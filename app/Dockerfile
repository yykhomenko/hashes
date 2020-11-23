FROM golang:latest
WORKDIR /go/src/github.com/cbi-sh/hashes/
#RUN go get -d -v golang.org/x/net/html
COPY ./main.go .
RUN GOOS=linux go build -o main .

#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=0 /go/src/github.com/cbi-sh/hashes/main .
EXPOSE 5000/tcp
CMD ["./main"]
