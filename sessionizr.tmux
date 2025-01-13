#!/bin/env bash

BIN_PATH=./bin/tmux-sessionizr
ARCHIVE_PATH=./bin/tmux-sessionizr.tar.gz

# Define GitHub repository and API endpoint for downloading the binary
REPO="vbargl/tmux-sessionizr"
API_URL="https://api.github.com/repos/$REPO/releases/tags"

function install_sessionizr {
  # Check if tmux-sessionizr is available in the PATH
  BIN="$(tmux show -gv @sessionizr-bin 2>/dev/null)"
  if command -v tmux-sessionizr >/dev/null 2>&1; then
    BIN=$(command -v tmux-sessionizr)
  fi

  # Get the latest release tag relative to current commit
  GIT_TAG=$(git describe --tags --abbrev=0)

  # Check if tmux-sessionizr is installed and up-to-date
  if [ -f "$BIN" ]; then
    INSTALLED_VERSION=$("$BIN" version)
    if [ "$INSTALLED_VERSION" == "dev" ]; then
      tmux set -g @sessionizr-bin "$BIN"
      exit 0
    fi

    if [ "$INSTALLED_VERSION" == "$GIT_TAG" ]; then
      tmux set -g @sessionizr-bin "$BIN"
      exit 0
    else
      echo "Updating tmux-sessionizr from $INSTALLED_VERSION to $GIT_TAG"
    fi
  fi

  mkdir -p ./bin

  # Get the release information
  RELEASE_INFO=$(curl -s "$API_URL/$GIT_TAG")

  # Get the OS and architecture
  OS=$(uname -s)
  ARCH=$(uname -m)

  # Build the download URL
  DOWNLOAD_URL="$(echo "$RELEASE_INFO" | grep -oP "\"browser_download_url\": \"\K[^\"]*${OS}_${ARCH}[^\"]+")"
  if [ -z "$DOWNLOAD_URL" ]; then
      echo "No binary found for platform $OS and architecture $ARCH."
      exit 1
  fi

  # Download the binary
  echo "Downloading tmux-sessionizr from $DOWNLOAD_URL"
  curl -L "$DOWNLOAD_URL" -o "$ARCHIVE_PATH"
  tar -xzf "$ARCHIVE_PATH" -C ./bin
  rm "$ARCHIVE_PATH"
  chmod +x "$BIN_PATH"

  tmux set-option -g @sessionizr-bin "$BIN_PATH"
}

install_sessionizr



# Bind the key to run 'pick'
# bind S run-shell './bin/tmux-sessionizr pick'

