FROM golang:alpine as build-env
RUN mkdir /main
WORKDIR /main
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ../.. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o main ./cmd/hashes

FROM scratch
COPY --from=build-env /main .
ENTRYPOINT ["/main"]
