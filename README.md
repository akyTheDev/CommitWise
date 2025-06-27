# CommitWise

A simple, fast, and powerful CLI tool that uses AI to generate Conventional Commit messages from a git diff.

## Overview

`commitwise` is a command-line utility that acts as a smart filter. You provide it with a git diff, and it returns a well-formatted commit message. It's designed to be a seamless part of your Git workflow by integrating with the pipe (|) operator.

## Features

- **AI-Powered**: Uses OpenAI's GPT models to generate high-quality commit messages.
- **Conventional Commits**: Adheres to the Conventional Commits standard for a clean and readable Git history.
- **Pipe-Based**: Follows the Unix philosophy—it does one thing well and works with other tools.

## Installation

### Prerequisites

- Go (version 1.18 or later)
- Git

### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/akyTheDev/CommitWise.git
   cd CommitWise
   ```
2. Build the binary:
   ```bash
   go build -o commitwise ./cmd/commitwise
   ```
3. Move the binary to your system's PATH:
   ```bash
   sudo mv commitwise /usr/local/bin/
   ```

### Configuration

The tool requires an OpenAI API key. Set it as an environment variable in your shell's configuration file (~/.zshrc, ~/.bashrc, etc.).# Add this line to your shell's config file

```bash
export OPENAI_API_KEY="sk-YourSecretApiKeyGoesHere"
```

Remember to restart your terminal or source the config file (e.g., `source ~/.zshrc`).

### Usage

`commitwise` is designed to be used with a pipe. You must provide a diff to its standard input.

#### Standard Workflow (Staged Changes)

This is the primary use case.

```bash
# Generate a message for all currently staged changes
git diff --staged | commitwise
```

#### Copying the Result

Pipe the output directly to your system's clipboard utility.

**On macOS**:

```bash
git diff --staged | commitwise | pbcopy
```

**On Linux**:

```bash
git diff --staged | commitwise | xclip -selection clipboard
```
