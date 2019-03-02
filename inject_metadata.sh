#!/bin/sh
set -e

sed "s|BUILD_VERSION|${0}|g" -i app-metadata.json
sed "s|COMMIT_SHA|${1}|g" -i app-metadata.json
