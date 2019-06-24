FROM golang:latest as builder
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
 
RUN CGO_ENABLED=0 GOOS=linux go build -o main . 
CMD ["/app/main"]
FROM alpine:latest AS production
# We have to copy the output from our
# builder stage to our production stage
COPY --from=builder /app .
# we can then kick off our newly compiled
# binary exectuable!!
CMD ["./main"]