FROM golang:1.9 as builder
WORKDIR /go/src/github.com/RossMerr/Caudex.Graph/cmd/caudex
COPY . /go/src/github.com/RossMerr/Caudex.Graph
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/RossMerr/Caudex.Graph/cmd/caudex/app .
EXPOSE 3000
CMD ["./app"]  