#!/bin/bash
set -eo pipefail

go build -v
./generator proto
