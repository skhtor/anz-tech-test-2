#!/bin/bash
set -e

echo "Running lint.."
golint .
echo "Lint complete!"
echo ""

echo "Running unit tests.."
go test -v
echo "Unit tests complete!"
