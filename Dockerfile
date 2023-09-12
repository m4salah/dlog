FROM golang:1.19-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o dlog ./cmd/dlog

FROM alpine as final
COPY --from=builder /app/dlog /app/dlog

ENTRYPOINT ["/app/dlog"]
CMD ["-bind", "0.0.0.0:3000", "-source", "/files"]
