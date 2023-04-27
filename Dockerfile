FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make git && make build

#================================

FROM alpine:latest

COPY --from=build /project/bin/app /bin/

CMD ["app"]