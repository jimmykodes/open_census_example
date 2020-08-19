FROM golang:1.14-buster as builder
WORKDIR /go/src/app
COPY . .
ENV GOBIN=/
RUN go install -v -mod=vendor ./...

FROM debian:buster
COPY --from=builder /app /app
CMD ["/app"]