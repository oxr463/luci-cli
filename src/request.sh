#!/bin/sh
set -e

request() {
  curl "${URL}" --data "${1}"
}
