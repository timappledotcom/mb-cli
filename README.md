# mb - Micro.blog Poster

A pretty, simple, and quick command line tool to post to Micro.blog using Charm tools. Note: This tool is currently supported on Linux only.

## Installation

### From Package

**Debian/Ubuntu (.deb):**
```bash
# Download the .deb from releases
sudo dpkg -i mb_*.deb
```

**Fedora/RHEL (.rpm):**
```bash
# Download the .rpm from releases
sudo rpm -i mb_*.rpm
```

**Arch Linux:**
```bash
# Download the .pkg.tar.zst from releases
sudo pacman -U mb_*.pkg.tar.zst
```

### From Source

```bash
git clone https://github.com/timappledotcom/mb-cli
cd mb-cli
go install
```

## Usage

### Setup
Run the command for the first time to set up your API keys interactively.

```bash
mb
```

You will need:
- Micro.blog App Token (Account > App Tokens)

### Posting
Interactive mode (opens a text area):
```bash
mb
```

Quick mode:
```bash
mb "Just setting up my new CLI posting tool! 🚀"
```

## Configuration
Configuration is stored in `~/.config/mb/config.json`.
