From golang:1.11-alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/cocoagaurav/httpHandler
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["httpHandler"]

EXPOSE 8080