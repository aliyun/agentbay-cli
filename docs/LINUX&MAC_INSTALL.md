# Agentbay Cloud Homebrew Tap

This is the official Homebrew tap for Agentbay Cloud, providing a convenient way to install the Agentbay command-line tool.

## Quick Installation

### Install via Tap (Recommended)

```bash
# 1. Add Agentbay Cloud's Homebrew tap
brew tap aliyun/agentbay

# 2. Install agentbay command-line tool
brew install agentbay

# 3. Verify installation
agentbay version
```

## Usage

After installation, you can use the following commands:

```bash
# Check version information
agentbay version

# View help information
agentbay --help

# Use Agentbay commands
agentbay [command] [options]
```

## Update and Uninstall

### Update to Latest Version

```bash
# update only agentbay
brew upgrade agentbay
```

### Uninstall

```bash
# Uninstall agentbay
brew uninstall agentbay

# Remove tap (optional)
brew untap aliyun/agentbay
```

## Troubleshooting

### Common Issues

1. **Installation Failed**
   ```bash
   # Update Homebrew
   brew update
   
   # Clean cache
   brew cleanup
   
   # Reinstall
   brew reinstall agentbay
   ```

2. **Network Issues (Especially for Chinese Users)**
   
   The Formula has been configured with Chinese mirror sources. If you still have issues:
   ```bash
   # Set Go proxy
   export GOPROXY=https://goproxy.cn,direct
   export GOSUMDB=sum.golang.google.cn
   
   # Reinstall
   brew reinstall agentbay
   ```

3. **Permission Issues**
   ```bash
   # Fix Homebrew permissions
   sudo chown -R $(whoami) $(brew --prefix)/*
   ```





