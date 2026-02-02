# CLIProxyAPI Plus

English | [Chinese](README_CN.md)

This is the Plus version of [CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI), adding support for third-party providers on top of the mainline project.

All third-party provider support is maintained by community contributors; CLIProxyAPI does not provide technical support. Please contact the corresponding community maintainer if you need assistance.

The Plus release stays in lockstep with the mainline features.

## Differences from the Mainline

- Added GitHub Copilot support (OAuth login), provided by [em4go](https://github.com/em4go/CLIProxyAPI/tree/feature/github-copilot-auth)
- Added Kiro (AWS CodeWhisperer) support (OAuth login), provided by [fuko2935](https://github.com/fuko2935/CLIProxyAPI/tree/feature/kiro-integration), [Ravens2121](https://github.com/Ravens2121/CLIProxyAPIPlus/)

## New Features (Plus Enhanced)

- **OAuth Web Authentication**: Browser-based OAuth login for Kiro with beautiful web UI
- **Rate Limiter**: Built-in request rate limiting to prevent API abuse
- **Background Token Refresh**: Automatic token refresh 10 minutes before expiration
- **Metrics & Monitoring**: Request metrics collection for monitoring and debugging
- **Device Fingerprint**: Device fingerprint generation for enhanced security
- **Cooldown Management**: Smart cooldown mechanism for API rate limits
- **Usage Checker**: Real-time usage monitoring and quota management
- **Model Converter**: Unified model name conversion across providers
- **UTF-8 Stream Processing**: Improved streaming response handling

## Kiro Authentication

### Web-based OAuth Login

Access the Kiro OAuth web interface at:

```
http://your-server:8080/v0/oauth/kiro
```

This provides a browser-based OAuth flow for Kiro (AWS CodeWhisperer) authentication with:
- AWS Builder ID login
- AWS Identity Center (IDC) login
- Token import from Kiro IDE

### CLI Authentication

#### AWS Builder ID (Default)

```bash
# Authorization Code Flow (recommended, automatic browser callback)
./CLIProxyAPI --kiro-aws-authcode

# Device Code Flow (manual code entry)
./CLIProxyAPI --kiro-aws-login
```

#### AWS Identity Center (IDC)

For enterprise users with AWS Identity Center:

```bash
# Authorization Code Flow (default, automatic browser callback)
./CLIProxyAPI --kiro-idc-login --kiro-idc-start-url https://d-xxx.awsapps.com/start

# Device Code Flow (manual code entry, for headless environments)
./CLIProxyAPI --kiro-idc-login --kiro-idc-start-url https://d-xxx.awsapps.com/start --kiro-idc-flow device

# Specify OIDC region (default: us-east-1)
./CLIProxyAPI --kiro-idc-login --kiro-idc-start-url https://d-xxx.awsapps.com/start --kiro-idc-region us-east-2
```

**IDC Options:**
| Option | Description | Default |
|--------|-------------|---------|
| `--kiro-idc-login` | Enable IDC login mode | - |
| `--kiro-idc-start-url` | Your IDC Start URL (required) | - |
| `--kiro-idc-region` | OIDC region for login and token refresh | `us-east-1` |
| `--kiro-idc-flow` | Login flow: `authcode` or `device` | `authcode` |

#### Google OAuth (Not Available)

> **Note:** Google/GitHub login is not available for third-party applications due to AWS Cognito redirect_uri restrictions. Please use AWS Builder ID or import your token from Kiro IDE.

```bash
./CLIProxyAPI --kiro-google-login
```

#### Import from Kiro IDE

If you've already logged in via Kiro IDE:

```bash
./CLIProxyAPI --kiro-import
```

## Quick Deployment with Docker

### One-Command Deployment

```bash
# Create deployment directory
mkdir -p ~/cli-proxy && cd ~/cli-proxy

# Create docker-compose.yml
cat > docker-compose.yml << 'EOF'
services:
  cli-proxy-api:
    image: eceasy/cli-proxy-api-plus:latest
    container_name: cli-proxy-api-plus
    ports:
      - "8317:8317"
    volumes:
      - ./config.yaml:/CLIProxyAPI/config.yaml
      - ./auths:/root/.cli-proxy-api
      - ./logs:/CLIProxyAPI/logs
    restart: unless-stopped
EOF

# Download example config
curl -o config.yaml https://raw.githubusercontent.com/router-for-me/CLIProxyAPIPlus/main/config.example.yaml

# Pull and start
docker compose pull && docker compose up -d
```

### Configuration

Edit `config.yaml` before starting:

```yaml
# Basic configuration example
server:
  port: 8317

# Add your provider configurations here
```

### Update to Latest Version

```bash
cd ~/cli-proxy
docker compose pull && docker compose up -d
```

## Contributing

This project only accepts pull requests that relate to third-party provider support. Any pull requests unrelated to third-party provider support will be rejected.

If you need to submit any non-third-party provider changes, please open them against the [mainline](https://github.com/router-for-me/CLIProxyAPI) repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
