FROM golang:1.22
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM busybox
COPY --from=0 /go/bin/app /app
