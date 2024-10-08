#!/bin/sh

usage() {
    echo "Usage: $0 -i input_file -o output_file"
    exit 1
}

while getopts "i:o:" opt; do
    case $opt in
        i) input_file="$OPTARG" ;;
        o) output_file="$OPTARG" ;;
        *) usage ;;
    esac
done

if [ -z "$input_file" ] || [ -z "$output_file" ]; then
    usage
fi



while IFS= read -r line; do
    number=$(echo "$line" | sed -n "s/.*'number': '\([^']*\)'.*/\1/p")
    name=$(echo "$line" | sed -n "s/.*'name': '\([^']*\)'.*/\1/p")
    script=$(echo "$line" | sed -n "s/.*'script': '\([^']*\)'.*/\1/p" | sed 's/\.py$/.go/')
    section=$(echo "$line" | sed -n "s/.*'section': '\([^']*\)'.*/\1/p")

    if [ -n "$number" ] && [ -n "$name" ] && [ -n "$script" ] && [ -n "$section" ]; then
        echo "    {" >> "$output_file"
        echo "        \"number\":  $number," >> "$output_file"
        echo "        \"name\":    \"$name\"," >> "$output_file"
        echo "        \"script\":  \"$script\"," >> "$output_file"
        echo "        \"section\": \"$section\"," >> "$output_file"
        echo "    }," >> "$output_file"
    fi
done < "$input_file"

sed -i '$ s/,$//' "$output_file"

echo "Conversion completed. Output saved to $output_file"
