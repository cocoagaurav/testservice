From golang:1.11-alpine

RUN apk update && apk add --no-cache git
RUN mkdir -p scrap

WORKDIR /go/src/github.com/cocoagaurav/scrap
COPY . ./scrap

RUN go get -d -v ./scrap/...
RUN go install ./scrap/...

CMD ["scrap"]


EXPOSE 8082