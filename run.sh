#!/bin/bash
cd rust/
cargo run --release
cd ../go/
go run ./...
cd ..

