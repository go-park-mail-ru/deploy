FROM golang:alpine

COPY . /project

WORKDIR /project

RUN apk add make && make build

CMD ["/project/app"]