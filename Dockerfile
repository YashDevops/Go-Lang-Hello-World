FROM golang:alpine as builder
LABEL name="yash kumar shah" maintainer="yashkumarshah22@gmail.com"
# git to download go lang dependencies
RUN apk update && apk add --no-cache git
COPY . .
RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o grab  .
ENTRYPOINT ["./grab"]
#FROM scratch
#Copy our static executable.
#COPY --from=builder /go/grab /go/grab 
#EXPOSE 8080
#ENTRYPOINT ["./go/grab"]
