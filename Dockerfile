FROM golang:buster

RUN mkdir /app && \
    go get github.com/aws/aws-sdk-go && \
    go get github.com/jinzhu/gorm && \
    go get github.com/gorilla/mux


ADD . /app
WORKDIR /app
RUN go build -o main.exe .
ENTRYPOINT ["/app/main.exe"]
