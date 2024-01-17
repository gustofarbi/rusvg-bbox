FROM rust:1.75.0-alpine3.19 as builder

RUN apk update && apk add -q musl-dev

WORKDIR /usr/src/app

COPY Cargo.toml Cargo.lock ./
COPY src ./src

RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/usr/src/app/target \
    cargo build --release --bin server \
    && mv target/release/server /tmp/server


FROM alpine:3.19.0

COPY --from=builder /tmp/server /usr/local/bin/server

CMD ["server"]
