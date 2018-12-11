From golang:1.11-alpine

RUN apk update && apk add --no-cache git
RUN mkdir -p post

WORKDIR /go/src/github.com/cocoagaurav/Post
COPY . ./post

RUN go get -d -v ./post/...
RUN go install ./post/...

CMD ["post"]


EXPOSE 8081