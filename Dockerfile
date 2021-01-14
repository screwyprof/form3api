FROM golang:1.14.13-stretch as build-env
RUN mkdir /form3api
WORKDIR /form3api
COPY . .
CMD make test