#!/bin/bash
# Download Ecommerce product dataset for benchmarking

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUTPUT_DIR="$SCRIPT_DIR/../../data"
mkdir -p "$OUTPUT_DIR"

echo "Ecommerce Product Dataset Download"
echo "================================="
echo ""
echo "Options:"
echo "1. Download Walmart sample (~1000 products) from GitHub"
echo "2. Use synthetic data generator (recommended for benchmarks)"
echo ""
read -p "Choose option (1/2): " choice

case $choice in
    1)
        echo "Downloading Walmart sample dataset..."
        cd "$OUTPUT_DIR"
        
        # Download Walmart products CSV
        wget -q --show-progress -O walmart_products.csv \
            "https://raw.githubusercontent.com/luminati-io/Walmart-dataset-samples/main/walmart-products.csv"
        
        echo "Downloaded to $OUTPUT_DIR/walmart_products.csv"
        
        # Also download Etsy for more variety
        wget -q --show-progress -O etsy_products.csv \
            "https://raw.githubusercontent.com/luminati-io/Etsy-dataset-sample/main/Etsy-dataset-sample.csv"
        
        echo "Downloaded to $OUTPUT_DIR/etsy_products.csv"
        ;;
    2)
        echo ""
        echo "Using synthetic data generator (recommended):"
        echo "  go run ./cmd/loader/main.go"
        echo ""
        echo "This generates unlimited realistic product data"
        ;;
    *)
        echo "Invalid option"
        exit 1
        ;;
esac

echo ""
echo "Done!"
