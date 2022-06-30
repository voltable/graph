FROM golang:1.18 as builder
WORKDIR /go/src/github.com/voltable/graph/cmd/caudex
COPY . /go/src/github.com/voltable/graph
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/voltable/graph/cmd/caudex/app .
EXPOSE 3000
CMD ["./app"]  