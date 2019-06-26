FROM golang:latest as builder

RUN mkdir /app 

ADD . /app/ 

WORKDIR /app

ARG POSTGRES_URI
ENV POSTGRES_URI=${POSTGRES_URI}
ARG API_KEY
ENV API_KEY=${API_KEY}

RUN CGO_ENABLED=0 GOOS=linux go build -o main . 

CMD ["/app/main"]

FROM alpine:latest AS production

COPY --from=builder /app .

ENTRYPOINT [ "./main" ]