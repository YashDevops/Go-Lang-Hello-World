FROM golang:alpine as builder
LABEL name="yash kumar shah" maintainer="yashkumarshah22@gmail.com"
RUN apk update && apk add --no-cache git && apk add gcc
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o grab  .
#ENTRYPOINT ["./grab"]

FROM scratch
COPY --from=builder /go/grab /
EXPOSE 8080
ENTRYPOINT ["/grab"]
