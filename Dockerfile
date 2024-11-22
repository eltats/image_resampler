FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o img-resampler ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8085

ENTRYPOINT ["./img-resampler"]

CMD ["-path-orig", "/tmp/img_orig", "-path-res", "/tmp/img_res", "-width", "200", "-height", "200"]
