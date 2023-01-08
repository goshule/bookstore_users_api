#base go image
# FROM golang:1.19-alpine as builder
# RUN mkdir /app
# COPY .  /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o bookstoreapi ./
# RUN chmod +x /app/bookstoreapi

#build a tiny docker image

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
#COPY --from=builder /app/bookstoreapi /app
# COPY ./creds.env ./app/
# COPY creds.env ./
# COPY bookstore-user-api /app/
# COPY bookstore-user-api ./
# #COPY ./bookstore-user-api /app
# RUN chmod +x /app/bookstore-user-api
COPY bookstore-user-api /app/   
#RUN chmod +x bookstore-user-api
CMD ["/app/bookstore-user-api"]