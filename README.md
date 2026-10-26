# silvergrab — SilverArch Package Manager

> Beautiful pacman wrapper with paru-style UI

![silvergrab](https://img.shields.io/badge/SilverArch-Package_Manager-00BFFF?style=flat)
![License](https://img.shields.io/badge/License-GPL--v3.0-green)
![Go](https://img.shields.io/badge/Go-1.22-00ADD8)

## Features

✨ **Beautiful CLI**
- Paru-style colored output
- ASCII art banner
- Clear package listings with numbering
- Colored confirmations (green/yellow/red)

🚀 **Fast & Simple**
- Thin wrapper around pacman
- Instant feedback
- Interactive confirmations

📦 **Operations**
- Install packages: `silvergrab -S <pkg>`
- Remove packages: `silvergrab -R <pkg>`
- Search packages: `silvergrab -Ss <query>`
- System upgrade: `silvergrab -Syu`

## Installation

```bash
cd ~/Projects/silvergrab
go build -o silvergrab ./cmd/
sudo cp silvergrab /usr/local/bin/
```

Or add to PATH:
```bash
export PATH="$PATH:$HOME/Projects/silvergrab"
```

## Usage

### Install Package(s)
```bash
sudo silvergrab -S firefox
sudo silvergrab -S firefox neovim git
```

### Remove Package(s)
```bash
sudo silvergrab -R vim
# asks: keep config files? [y/N]
# asks: proceed? [y/N]
```

### Search Packages
```bash
silvergrab -Ss neovim
silvergrab -Ss "text editor"
```

### System Upgrade
```bash
sudo silvergrab -Syu
# syncs all repos and upgrades all packages
```

### Help & Version
```bash
silvergrab --help
silvergrab --version
```

## Options

| Option | Alias | Description |
|--------|-------|-------------|
| `-S` | `--sync` | Install/sync packages |
| `-R` | `--remove` | Remove packages |
| `-Ss` | `--search` | Search repositories |
| `-Syu` | `--sync-upgrade` | Sync + full system upgrade |
| `-h` | `--help` | Show help message |
| `-v` | `--version` | Show version |

## Requirements

- Go 1.22+
- pacman (ArchLinux/SilverArch)
- Root access (for install/remove/upgrade)

## Dependencies

- `github.com/charmbracelet/lipgloss` — terminal styling

## Color Scheme

| Color | Usage |
|-------|-------|
| 🔵 Cyan | Headers + titles |
| 🟢 Green | Success messages + package numbers |
| 🟡 Yellow | Warnings (destructive operations) |
| 🔴 Red | Errors |
| ⚪ Silver | Logo + accents |

## Project Structure

```
silvergrab/
├── cmd/
│   └── main.go      # CLI entry point
├── go.mod           # Go module definition
├── go.sum           # Dependencies lock file
├── README.md        # This file
└── silvergrab       # Compiled binary
```

## Development

Build:
```bash
go build -o silvergrab ./cmd/
```

Test:
```bash
./silvergrab --help
./silvergrab --version
silvergrab -Ss firefox
```

## Author

**Silvercore (s1lverarch)**
- Email: s1lv3rcore@proton.me
- Website: https://s1lverarch.github.io/linuxdistro/
- GitHub: https://github.com/s1lverarch

## License

GPL-v3.0 — See LICENSE file

## Related Projects

- [SilverSaphire](https://github.com/s1lverarch/SilverSaphire) — Installer
- [silverflux](https://github.com/s1lverarch/silverflux) — Sync + mirrors
- [silverpurge](https://github.com/s1lverarch/silverpurge) — System cleanup

---

**SilverArch Linux Project** — Made With Love In Tunisia
