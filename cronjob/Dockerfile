From golang:1.11-alpine

RUN apk update && apk add --no-cache git

RUN mkdir -p cron

WORKDIR /go/src/github.com/cocoagaurav/cronjob

COPY . cron

RUN go get -d -v ./cron/...

RUN go install -v ./cron/...

CMD ["cron"]

EXPOSE 8085