FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/example/calculator-server/
COPY . .

RUN go get .

RUN CGO_ENABLED=0 go build -o /go/bin/calculator-server

FROM scratch
COPY --from=builder /go/bin/calculator-server /go/bin/calculator-server

EXPOSE 8080
ENTRYPOINT ["/go/bin/calculator-server"]
