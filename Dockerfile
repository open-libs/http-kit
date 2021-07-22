#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/src/app/main

#final stage
FROM alpine:latest
ENV app_name=http-kit
RUN apk --no-cache add ca-certificates && \
    mkdir /conf /log /app
COPY --from=builder /go/src/app/main /app/${app_name}
ENTRYPOINT /app/${app_name}
EXPOSE 80
