#!/usr/bin/env bash
set -euo pipefail

REPO="defenseclaw/defenseclaw"
BINARY="defenseclaw"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "${ARCH}" in
    x86_64)  ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    arm64)   ARCH="arm64" ;;
    *)       echo "Unsupported architecture: ${ARCH}"; exit 1 ;;
esac

echo "Detected: ${OS}/${ARCH}"
echo "Installing ${BINARY} to ${INSTALL_DIR}..."

LATEST=$(curl -sSf "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
URL="https://github.com/${REPO}/releases/download/${LATEST}/${BINARY}-${OS}-${ARCH}"

curl -sSfL "${URL}" -o "${INSTALL_DIR}/${BINARY}"
chmod +x "${INSTALL_DIR}/${BINARY}"

echo "${BINARY} ${LATEST} installed to ${INSTALL_DIR}/${BINARY}"
echo ""
echo "Run 'defenseclaw init' to get started."
