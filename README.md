# calc-cli

A simple command-line calculator written in Go.

It evaluates arithmetic expressions (addition, subtraction, multiplication, division, parentheses and decimals) and provides extra commands for logarithms and square roots.

```sh
calc "2 + 3 * 4"        # 14
calc "(10 - 2) / 4"     # 2
calc log 8 -b 2         # 3
calc sqrt 16            # 4
```

## Installation

Prebuilt binaries for Windows, Linux and macOS (amd64 and arm64) are published on the [Releases page](https://github.com/YasserHage/calc-cli/releases).

Archives follow the pattern `calc-cli_<version>_<os>_<arch>`, and each one contains a single `calc` binary (`calc.exe` on Windows).

| OS      | Intel / AMD (x86-64)                  | ARM (arm64)                           |
| ------- | ------------------------------------- | ------------------------------------- |
| Windows | `calc-cli_<version>_windows_amd64.zip`   | `calc-cli_<version>_windows_arm64.zip`   |
| Linux   | `calc-cli_<version>_linux_amd64.tar.gz`  | `calc-cli_<version>_linux_arm64.tar.gz`  |
| macOS   | `calc-cli_<version>_darwin_amd64.tar.gz` | `calc-cli_<version>_darwin_arm64.tar.gz` |

The examples below use version `1.0.0`; replace it with the [latest release](https://github.com/YasserHage/calc-cli/releases/latest) if a newer one is available.

### Windows

1. Download `calc-cli_1.0.0_windows_amd64.zip` (or `windows_arm64` on ARM devices) from the Releases page.
2. Extract the zip and move `calc.exe` to a folder of your choice, e.g. `C:\Program Files\calc`.
3. Add that folder to your `PATH`:
   - Open **Start** → search for **Edit the system environment variables** → **Environment Variables...**
   - Under **User variables**, select `Path` → **Edit** → **New**, and add `C:\Program Files\calc`.
4. Open a new terminal and check the installation:

   ```powershell
   calc --help
   ```

Or do it all from PowerShell:

```powershell
$version = "1.0.0"
$dir = "$env:LOCALAPPDATA\Programs\calc"
New-Item -ItemType Directory -Force $dir | Out-Null
Invoke-WebRequest "https://github.com/YasserHage/calc-cli/releases/download/v$version/calc-cli_${version}_windows_amd64.zip" -OutFile "$env:TEMP\calc.zip"
Expand-Archive "$env:TEMP\calc.zip" -DestinationPath $dir -Force
[Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "User") + ";$dir", "User")
```

Restart your terminal afterwards so the updated `PATH` is picked up.

### Linux

```sh
VERSION=1.0.0
ARCH=amd64   # use arm64 on ARM machines (check with: uname -m)

curl -LO "https://github.com/YasserHage/calc-cli/releases/download/v${VERSION}/calc-cli_${VERSION}_linux_${ARCH}.tar.gz"
tar -xzf "calc-cli_${VERSION}_linux_${ARCH}.tar.gz" calc
sudo install -m 755 calc /usr/local/bin/calc

calc --help
```

Without `sudo`, install it to `~/.local/bin` instead (make sure that directory is on your `PATH`):

```sh
mkdir -p ~/.local/bin && install -m 755 calc ~/.local/bin/calc
```

### macOS

```sh
VERSION=1.0.0
ARCH=arm64   # Apple Silicon (M1/M2/M3...); use amd64 on Intel Macs

curl -LO "https://github.com/YasserHage/calc-cli/releases/download/v${VERSION}/calc-cli_${VERSION}_darwin_${ARCH}.tar.gz"
tar -xzf "calc-cli_${VERSION}_darwin_${ARCH}.tar.gz" calc
sudo install -m 755 calc /usr/local/bin/calc

calc --help
```

The binary is not signed, so macOS Gatekeeper may block it the first time it runs. If you see a message saying the developer cannot be verified, remove the quarantine attribute:

```sh
sudo xattr -d com.apple.quarantine /usr/local/bin/calc
```

### Verifying the download (optional)

Each release includes a `calc-cli_<version>_checksums.txt` file with SHA-256 checksums:

```sh
# Linux
sha256sum --ignore-missing -c calc-cli_1.0.0_checksums.txt
# macOS
shasum -a 256 --ignore-missing -c calc-cli_1.0.0_checksums.txt
```

```powershell
# Windows: compare the output with the matching line in the checksums file
Get-FileHash .\calc-cli_1.0.0_windows_amd64.zip -Algorithm SHA256
```

### Using `go install` (any OS)

If you have Go 1.26 or newer installed:

```sh
go install github.com/YasserHage/calc-cli@latest
```

The binary is installed to `$(go env GOPATH)/bin` and is named `calc-cli`. Make sure that directory is on your `PATH`.

### Building from source

```sh
git clone https://github.com/YasserHage/calc-cli.git
cd calc-cli
go build -o calc .        # on Windows: go build -o calc.exe .
```

## Usage

### Expressions

Pass the expression as a quoted string so your shell does not interpret special characters such as `*` or `(`:

```sh
calc "2 + 3 * 4"         # 14
calc "(10 - 2) / 4"      # 2
calc "3.5 * (2 + 1)"     # 10.5
```

Supported: `+`, `-`, `*`, `/`, parentheses, decimal and negative numbers (e.g. `calc "2 * -3"`).

### `log`

Calculates the logarithm of a number. Uses base 10 by default; set a custom base with `-b` / `--base`.

```sh
calc log 100             # 2
calc log 8 -b 2          # 3
```

The number must be greater than 0, and the base must be greater than 0 and different from 1.

### `sqrt`

Calculates the square root of a non-negative number.

```sh
calc sqrt 9              # 3
calc sqrt 2              # 1.4142135623730951
```

Run `calc --help` or `calc <command> --help` for more details.

## Development

```sh
go test ./...            # run the tests
go vet ./...             # static checks
gofmt -l .               # list unformatted files
```

Releases are built with [GoReleaser](https://goreleaser.com/) by the `Release` GitHub Actions workflow whenever a tag is pushed.

## License

See [LICENSE](LICENSE).
