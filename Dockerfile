FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make && make build

FROM alpine:latest

COPY --from=build /project/app /bin/

CMD ["app"]