package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"os/user"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// Command-line flag for checking updates
	checkUpdates := flag.Bool("updates", false, "Check for system updates")
	flag.Parse()

	// Create color functions
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	//	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	// Get user info
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get host info
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}

	// Get memory info
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	// Get disk info
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	// Get IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	var ipAddress string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddress = ipnet.IP.String()
				break
			}
		}
	}

	// Get Linux distribution
	distro := "N/A"
	catCmd := exec.Command("cat", "/etc/os-release")
	out, err := catCmd.Output()
	if err == nil {
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				distro = strings.Trim(line[13:], `"`)
				break
			}
		}
	}

	// Print information
	fmt.Println(green("==OS Info v0.1=="))
	fmt.Println(cyan("--- 🖥️ System Information ---"))
	fmt.Printf("👤 User: %s\n", green(currentUser.Username))
	fmt.Printf("🐧 Distro: %s\n", green(distro))
	fmt.Printf("🌐 Hostname: %s\n", green(hostInfo.Hostname))
	fmt.Printf("🧠 Kernel: %s\n", green(hostInfo.KernelVersion))
	fmt.Printf("📡 IP Address: %s\n", green(ipAddress))

	fmt.Println(cyan("\n--- 📊 Usage Information ---"))
	fmt.Printf("💾 Memory: %s (%.2f%%)\n", yellow(fmt.Sprintf("%d/%d MB", memInfo.Used/1024/1024, memInfo.Total/1024/1024)), memInfo.UsedPercent)
	fmt.Printf("💽 Disk: %s (%.2f%%)\n", yellow(fmt.Sprintf("%d/%d GB", diskInfo.Used/1024/1024/1024, diskInfo.Total/1024/1024/1024)), diskInfo.UsedPercent)

	// Check for updates if the flag is provided
	if *checkUpdates {
		fmt.Println(cyan("\n--- ⬆️ Update Status ---"))
		fmt.Println("📦 Checking for updates... (This may require your password)")
		updates := "N/A"
		aptCheck := exec.Command("sudo", "apt-get", "update")
		if err := aptCheck.Run(); err == nil {
			aptSimulate := exec.Command("sudo", "apt-get", "upgrade", "-s")
			upgradable, err := aptSimulate.Output()
			if err == nil {
				count := strings.Count(string(upgradable), "Inst")
				if count > 0 {
					updates = fmt.Sprintf("%d packages can be upgraded.", count)
				} else {
					updates = "System is up-to-date."
				}
			}
		}
		fmt.Printf("Status: %s\n", magenta(updates))
	}
}
