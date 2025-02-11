FROM docker.io/library/golang:1.21-alpine3.17 AS builder

RUN apk --no-cache add build-base bash curl
RUN cd / && curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs -o rust.sh && sh rust.sh -y
RUN mkdir /poseidongold
ADD . /poseidongold
WORKDIR /poseidongold/rust
RUN source $HOME/.cargo/env && rustup override set nightly && cargo build --release
