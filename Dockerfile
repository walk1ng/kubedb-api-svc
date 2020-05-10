FROM golang:1.11 as build

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /go/cache

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch as prod
COPY --from=build /go/release/app /

ENTRYPOINT [ "/app" ]
