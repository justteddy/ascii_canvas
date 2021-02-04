FROM golang:1.15-alpine as build

WORKDIR $GOPATH/src/canvas
COPY . .
RUN go build -o /go/bin/canvas .

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
COPY --from=build /go/bin/canvas /bin/

ENTRYPOINT ["canvas"]
