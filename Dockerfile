FROM golang:1.22 as builder
WORKDIR /build
COPY  go.mod .
COPY  go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/car-factory ./cmd/app/main.go

FROM alpine:3.18.6
COPY --from=builder /build/bin/car-factory  /
COPY .env /
ENTRYPOINT ["/car-factory"]