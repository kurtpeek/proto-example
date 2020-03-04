#!/bin/bash

cd proto
protoc --go_out=paths=source_relative:../gen/go author/author.proto
protoc --go_out=paths=source_relative:../gen/go book/book.proto