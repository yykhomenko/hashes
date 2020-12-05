FROM golang:alpine
WORKDIR /go/src/github.com/cbi-sh/hashes/
COPY ./cmd/node/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0 ./go/src/github.com/cbi-sh/hashes/main .
CMD ["./main"]
