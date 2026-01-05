# App Downloader

Independent tool for downloading 1Panel App Store resources.

## Build

Build for all supported platforms:

```bash
make all
```

Build for specific platform:

```bash
make build-mac          # macOS (arm64)
make build-linux-amd64  # Linux amd64
make build-linux-arm64  # Linux arm64
```

Artifacts will be placed in the `bin/` directory.

## Usage

```bash
./bin/app-downloader-[platform]-[arch] [flags]
```

### Flags

- `-repo`: App Store Repository URL (Default: `https://apps-assets.fit2cloud.com`)
- `-mode`: Mode, `stable` or `dev` (Default: `stable`)

### Example

```bash
# Default
./bin/app-downloader-darwin-arm64

# Custom repo and mode
./bin/app-downloader-darwin-arm64 -repo "https://apps-assets.fit2cloud.com" -mode "dev"
```
