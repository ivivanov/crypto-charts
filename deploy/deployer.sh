#!/bin/bash
set -euo pipefail

# Check if both arguments are provided
if [ $# -ne 1 ]; then
    echo "Error: One argument(s) are required."
    echo "Usage: $0 <version>"
    exit 1
fi


# Assign arguments to variables
version=$1

# Validate arguments are not empty
if [ -z "$version" ]; then
    echo "Error: Argument(s) must not be empty."
    exit 1
fi

export VERSION=$version

docker compose --file './deploy/docker-compose.yml' up --detach --build
