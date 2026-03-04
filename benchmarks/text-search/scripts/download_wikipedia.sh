#!/bin/bash
# Download Wikipedia dataset for benchmarking
# Uses Simple Wikipedia from HuggingFace - ~770K articles, 87MB

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUTPUT_DIR="$SCRIPT_DIR/../../data"
mkdir -p "$OUTPUT_DIR"

echo "Downloading Simple Wikipedia dataset from HuggingFace..."
echo "This will download ~87MB (770K articles)"
echo ""

cd "$OUTPUT_DIR"

# Download the correct parquet file
wget -O wikipedia.parquet \
    "https://huggingface.co/datasets/rahular/simple-wikipedia/resolve/main/data/train-00000-of-00001-090b52ccb189d47a.parquet"

echo ""
echo "Download complete!"
echo "File: $OUTPUT_DIR/wikipedia.parquet"
echo ""
echo "To load this data in Go, you'll need a parquet library"
echo "Or convert to JSON using: pip install pandas pyarrow"
echo "  python -c \"import pandas as pd; df = pd.read_parquet('wikipedia.parquet'); df.to_json('wikipedia.json', orient='records')\""
