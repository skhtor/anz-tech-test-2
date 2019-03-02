#!/bin/sh
set -e

sed "s|BUILD_VERSION|$1|g" -i app_metadata.json
sed "s|COMMIT_SHA|$2|g" -i app_metadata.json
