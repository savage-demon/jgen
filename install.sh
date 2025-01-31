#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

REPO="savage-demon/jgen"
LATEST_VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [[ -z "$LATEST_VERSION" ]]; then
  echo "Error: Unable to fetch the latest version."
  exit 1
fi

VERSION_NO_PREFIX=$(echo "$LATEST_VERSION" | sed 's/^v//')

BINARY_NAME="jgen"
URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/${BINARY_NAME}_${VERSION_NO_PREFIX}_${OS}_${ARCH}.tar.gz"

HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" $URL)
if [[ "$HTTP_STATUS" -ge 400 ]]; then
  echo "Error: Failed to download $URL. HTTP status $HTTP_STATUS."
  exit 1
fi

echo "Downloading $BINARY_NAME version $LATEST_VERSION from $URL"
curl -L $URL -o /tmp/$BINARY_NAME.tar.gz

echo "Installing $BINARY_NAME..."
sudo tar -xzf /tmp/$BINARY_NAME.tar.gz -C /usr/local/bin || {
  echo "Error: Failed to extract the archive."
  exit 1
}
sudo chmod +x /usr/local/bin/$BINARY_NAME

echo "$BINARY_NAME $LATEST_VERSION installed successfully!"
