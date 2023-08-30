FROM ubuntu:22.04

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  pdftojpeg

RUN apt-get update && apt-get install -y libc6

COPY ./build .
RUN echo $(ls -la)
EXPOSE 16080

CMD ["./main"]