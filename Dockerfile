FROM golang:latest as builder


RUN mkdir /app

ADD . /app/ 

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main . 

CMD ["/app/main"]

FROM alpine:latest AS production

COPY --from=builder /app .

ENTRYPOINT [ "./main" ]