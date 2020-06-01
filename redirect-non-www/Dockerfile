FROM golang:1.10-alpine as builder
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED="0"
WORKDIR /go/src/app
COPY . .
RUN go build -o redirect .

FROM scratch
COPY --from=builder /go/src/app/redirect /bin/redirect
EXPOSE 80
CMD ["redirect"]