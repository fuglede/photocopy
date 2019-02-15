FROM golang:1.11.5-alpine
RUN apk --update add imagemagick
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 8000
RUN go build -o main .
CMD ["/app/main"]
