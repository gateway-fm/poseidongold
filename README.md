## Build the Rust library

```
$ cd rust
$ rustup override set nightly-x86_64-unknown-linux-gnu
$ cargo build --release
```

## Test Go code

```
$ cd go
$ cp ../rust/target/release/librustposeidongold.a .
$ go test
```

## License
Apache License, Version 2.0 [LICENSE](LICENSE)