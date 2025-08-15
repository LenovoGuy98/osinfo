
# OS Info Script

This Go script provides a colorful and organized overview of your Linux system information. It displays details about your operating system, hardware, and network, and includes an optional feature to check for system updates.

## Features

- **System Information**: Displays your username, Linux distribution, hostname, kernel version, and IP address.
- **Usage Statistics**: Shows current memory and disk usage percentages and values.
- **Update Checker**: Includes an optional command-line flag to check for pending updates on Debian-based systems (like Ubuntu).
- **Colorful Output**: Uses colors and emojis to present the information in a user-friendly and readable format.

## Dependencies

This script uses the following Go libraries:

- `github.com/shirou/gopsutil`: To gather system information like memory, disk, and host details.
- `github.com/fatih/color`: To add color to the terminal output.

## Installation

1. **Install Go**: Make sure you have Go installed on your system.
2. **Install Dependencies**: Run the following commands to install the necessary libraries:
   ```bash
   go get github.com/shirou/gopsutil/v3/host
   go get github.com/shirou/gopsutil/v3/mem
   go get github.com/shirou/gopsutil/v3/disk
   go get github.com/shirou/gopsutil/v3/net
   go get github.com/fatih/color
   ```

## Usage

You can run the script in two ways:

1.  **Without Update Check**: To display the system information without checking for updates, run:
    ```bash
    go run osinfo.go
    ```

2.  **With Update Check**: To include the update check, use the `-updates` flag. This will prompt for your password to run `apt-get`.
    ```bash
    go run osinfo.go -updates
    ```

## Example Output

```
--- 🖥️ System Information ---
👤 User: your-username
🐧 Distro: Ubuntu 22.04.1 LTS
🌐 Hostname: your-hostname
🧠 Kernel: 5.15.0-48-generic
📡 IP Address: 192.168.1.10

--- 📊 Usage Information ---
💾 Memory: 4096/8192 MB (50.00%)
💽 Disk: 100/250 GB (40.00%)

--- ⬆️ Update Status ---
📦 Checking for updates... (This may require your password)
Status: 12 packages can be upgraded.
```
