FROM golang:1.20-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  pdftojpeg

WORKDIR /app
COPY . .
RUN echo $(ls)
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /main ./cmd/main.go
COPY ./config/ /config
EXPOSE 16080

CMD ["./main"]