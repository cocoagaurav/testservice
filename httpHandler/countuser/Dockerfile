From golang:onbuild

WORKDIR /go/src/github.com/cocoagaurav/httpHandler/countuser
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000