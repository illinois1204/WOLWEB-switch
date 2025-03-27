FROM golang:1.23.7-alpine3.21 AS builder

WORKDIR /workshop
COPY . .
RUN go mod download
RUN go build -o ./build/wolweb ./main.go


FROM alpine:3.21

RUN mkdir -p /app/store
COPY --from=builder ["/workshop/build/wolweb", "/app/"]
COPY --from=builder ["/workshop/views", "/app/views"]
COPY --from=builder ["/workshop/public", "/app/public"]

EXPOSE 80
WORKDIR /app
CMD ["./wolweb"]
