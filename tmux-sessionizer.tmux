#!/bin/env bash

TMUXSESSIONIZR_EXE=tmux-sessionizr

if ! command -v "$TMUXSESSIONIZR_EXE" >/dev/null; then
  tmux display-message "tmux-sessionizr not found"
fi

"$TMUXSESSIONIZR_EXE"