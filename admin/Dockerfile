
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/admin
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/admin /admin
# COPY --from=static . .
ENTRYPOINT ./admin
LABEL Name=app Version=0.0.1
EXPOSE 9000
