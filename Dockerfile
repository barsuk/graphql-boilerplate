FROM golang:1.13.0-alpine3.10

WORKDIR /go/src/graphql-boilerplate
COPY . .

RUN go version
RUN go build

CMD ./graphql-boilerplate

EXPOSE 4444
