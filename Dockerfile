FROM golang:latest as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app 
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app/
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]