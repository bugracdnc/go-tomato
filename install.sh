#!/bin/bash

# Define paths
EXECUTABLE_NAME="go-tomato"
INSTALL_PATH="$HOME/.local/bin/tomato"

# Remove previous build if it exists
[ -f "$EXECUTABLE_NAME" ] && rm "$EXECUTABLE_NAME"
[ -f "$INSTALL_PATH" ] && rm "$INSTALL_PATH"

# Build the executable
go build -o "$EXECUTABLE_NAME" "$EXECUTABLE_NAME"

# Copy it to ~/.local/bin/tomato
cp "$EXECUTABLE_NAME" "$INSTALL_PATH"

echo "Build complete and installed to $INSTALL_PATH"
