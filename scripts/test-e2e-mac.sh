#!/usr/bin/env bash
set -euo pipefail

echo "=== DefenseClaw E2E Test — macOS ==="

BINARY="./defenseclaw"

if [ ! -f "${BINARY}" ]; then
    echo "Binary not found. Run 'make build-darwin-arm64' first."
    exit 1
fi

echo "--- Init ---"
rm -rf ~/.defenseclaw
${BINARY} init

echo "--- Scan clean skill ---"
${BINARY} scan skill ./test/fixtures/skills/clean-skill/

echo "--- Scan malicious skill ---"
${BINARY} scan skill ./test/fixtures/skills/malicious-skill/ || true

echo "--- Scan AIBOM ---"
${BINARY} scan aibom .

echo "=== E2E tests complete ==="
