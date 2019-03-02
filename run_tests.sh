#!/bin/sh
set -e

echo "Running lint.."
golint .
echo "Lint complete!"
echo ""

echo "Running unit tests.."
export COMMIT_SHA="abc123"
export BUILD_VERSION="1.0"
go test -v
echo "Unit tests complete!"
