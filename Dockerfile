FROM golang:1.11.5-alpine as builder
WORKDIR /build
COPY . . 
RUN go build -o main .

FROM alpine
EXPOSE 8000
WORKDIR /root
RUN apk --update add imagemagick
COPY --from=builder /build/main .
COPY tmpl tmpl
CMD ["./main"]
