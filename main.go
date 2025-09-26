package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"os/user"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	// Initialize GTK.
	gtk.Init(nil)

	// Create a new toplevel window.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("OS Info")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new box container.
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(box)

	// Get OS info
	info := getOsInfo()

	// Create a label and add it to the box.
	label, err := gtk.LabelNew(info)
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	box.Add(label)

	// Set the default window size.
	win.SetDefaultSize(300, 200)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin the GTK main loop.
	gtk.Main()
}

func getOsInfo() string {
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
				distro = strings.Trim(line[13:], "`")
				break
			}
		}
	}

	// Format information
	return fmt.Sprintf(
		"--- System Information ---\n"+
		"User: %s\n"+
		"Distro: %s\n"+
		"Hostname: %s\n"+
		"Kernel: %s\n"+
		"IP Address: %s\n\n"+
		"--- Usage Information ---\n"+
		"Memory: %d/%d MB (%.2f%%)\n"+
		"Disk: %d/%d GB (%.2f%%)",
		currentUser.Username,
		distro,
		hostInfo.Hostname,
		hostInfo.KernelVersion,
		ipAddress,
		memInfo.Used/1024/1024, memInfo.Total/1024/1024, memInfo.UsedPercent,
		diskInfo.Used/1024/1024/1024, diskInfo.Total/1024/1024/1024, diskInfo.UsedPercent,
	)
}
