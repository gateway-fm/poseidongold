## Build the Rust library

```
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
$ cd rust
$ rustup override set nightly
$ cargo build --release
```


## Test Go code

```
$ cd go
$ cp ../rust/target/release/librustposeidongold.a .
$ go test
```

## Build for Alpine Linux

This is needed to build [xlayer-erigon](https://github.com/okx/xlayer-erigon).

```
$ docker run -v .:/poseidon -it docker.io/library/golang:1.22-alpine3.20 sh
# *** in the container:
# apk add curl
# apk add build-base
# cd /poseidon/rust
# curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
# . "$HOME/.cargo/env"
# rustup override set nightly
# cargo build --release
# cp target/release/librustposeidongold.a ../go/
# cd ../go
# go test
```

## License

Apache License, Version 2.0 [LICENSE](LICENSE)