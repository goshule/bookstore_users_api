#base go image
FROM golang:1.19-alpine as builder
RUN mkdir /app
COPY .  /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o bookstoreapi ./
RUN chmod +x /app/bookstoreapi

#build a tiny docker image

FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/bookstoreapi /app
COPY --from=builder /app/creds.env /app
CMD ["/app/bookstoreapi"]