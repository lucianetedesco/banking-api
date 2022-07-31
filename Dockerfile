FROM golang:1.18-alpine as builder

RUN mkdir /builder

COPY . /builder/

WORKDIR /builder

RUN go build -o /builder

FROM alpine:3.14

COPY --from=builder /builder/banking-api .

EXPOSE 8080

CMD ["/banking-api"]
