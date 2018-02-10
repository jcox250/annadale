FROM golang:1.9

WORKDIR /go/src/tech-test-jc
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["tech-test-jc"]
