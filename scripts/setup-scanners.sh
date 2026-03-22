#!/usr/bin/env bash
set -euo pipefail

echo "Installing DefenseClaw scanner dependencies..."
echo ""

pip install --upgrade pip

echo "Installing skill-scanner..."
pip install cisco-ai-skill-scanner

echo "Installing mcp-scanner..."
pip install mcp-scanner

echo "Installing aibom..."
pip install aibom

echo ""
echo "Scanner dependencies installed."
echo "Verify with: skill-scanner --version && mcp-scanner --version && aibom --version"
