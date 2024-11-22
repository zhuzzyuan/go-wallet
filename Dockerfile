FROM golang:1.22.1-alpine3.19
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
EXPOSE 8000
CMD ./main