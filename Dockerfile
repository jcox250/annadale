FROM golang:1.9

WORKDIR /go/src/github.com/jcox250/annadale
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["annadale"]
