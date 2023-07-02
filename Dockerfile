FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . /app
RUN go build -o main .

FROM alpine:3
WORKDIR /app/
COPY --from=builder /app/main /app
CMD /app/main