FROM golang:latest as builder
WORKDIR /
COPY . .

ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main .

EXPOSE 8080
CMD ["./main"]