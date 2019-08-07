#!/bin/bash

if [ -n "$(gofmt -l .)" ]; then
  echo "Go code is not formatted:"
  gofmt -d .
  gofmt -s -w .
  exit 1
fi
