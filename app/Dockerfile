FROM golang:latest
WORKDIR /go/src/github.com/cbi-sh/hashes/
COPY ./main.go .
RUN go build -o main .

#FROM alpine:latest
#COPY --from=0 ./go/src/github.com/cbi-sh/hashes/main .
EXPOSE 5000/tcp
CMD ["./main"]
