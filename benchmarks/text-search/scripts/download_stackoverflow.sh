#!/bin/bash
# Download Stack Overflow Q&A dataset for benchmarking
# Uses Stack Exchange dump from Archive.org

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUTPUT_DIR="$SCRIPT_DIR/../../data"
mkdir -p "$OUTPUT_DIR"

echo "Downloading Stack Overflow dataset from Archive.org..."
echo "WARNING: This is ~12GB compressed"
echo ""

# Option 1: Just Ask Ubuntu (smaller, ~1GB)
# Option 2: Full Stack Overflow (~12GB)
# Option 3: Use synthetic data (recommended)

echo "Options:"
echo "1. Download Ask Ubuntu (~1GB) - good for testing"
echo "2. Download Stack Overflow (~12GB) - full dataset"
echo "3. Use synthetic data (recommended)"
echo ""
read -p "Choose option (1/2/3): " choice

case $choice in
    1)
        echo "Downloading Ask Ubuntu..."
        cd "$OUTPUT_DIR"
        wget -q --show-progress -O askubuntu.7z \
            "https://archive.org/download/stackexchange/askubuntu.com.7z"
        7z x askubuntu.7z -o"askubuntu"
        echo "Downloaded to $OUTPUT_DIR/askubuntu/"
        ;;
    2)
        echo "Downloading Stack Overflow (this will take a while)..."
        cd "$OUTPUT_DIR"
        wget -q --show-progress -O stackoverflow.7z \
            "https://archive.org/download/stackexchange/stackoverflow.com.7z"
        7z x stackoverflow.7z -o"stackoverflow"
        echo "Downloaded to $OUTPUT_DIR/stackoverflow/"
        ;;
    3)
        echo "Using synthetic data generator instead..."
        echo "Run: go run ./cmd/loader/main.go"
        ;;
    *)
        echo "Invalid option"
        exit 1
        ;;
esac

echo ""
echo "Done!"
