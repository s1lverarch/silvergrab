// silvergrab — SilverArch Package Manager
// Author: Silvercore (s1lverarch) <s1lv3rcore@proton.me>
// License: GPL-v3.0

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	colorSilver = lipgloss.Color("#C0C0C0")
	colorBlue   = lipgloss.Color("#00BFFF")
	colorCyan   = lipgloss.Color("#00FFD0")
	colorGreen  = lipgloss.Color("#00FF7F")
	colorYellow = lipgloss.Color("#FFD700")
	colorRed    = lipgloss.Color("#FF4466")
	colorDim    = lipgloss.Color("#444444")
	colorDim2   = lipgloss.Color("#666666")

	styleLogo    = lipgloss.NewStyle().Foreground(colorSilver).Bold(true)
	styleHeader  = lipgloss.NewStyle().Foreground(colorCyan).Bold(true)
	styleOK      = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	styleWarn    = lipgloss.NewStyle().Foreground(colorYellow).Bold(true)
	styleErr     = lipgloss.NewStyle().Foreground(colorRed).Bold(true)
	styleInfo    = lipgloss.NewStyle().Foreground(colorBlue)
	styleNum     = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	styleDim     = lipgloss.NewStyle().Foreground(colorDim2)
	styleSubtl   = lipgloss.NewStyle().Foreground(colorDim)
)

const (
	version = "0.1.0"
	logo    = `
 ██████╗██████╗
 ██╔════╝██╔══██╗
 ███████╗█████╔╝
 ╚════██║██╔══██╗
 ███████║██████╔╝
 ╚══════╝╚═════╝
`
)

func main() {
	if len(os.Args) < 2 {
		showBanner()
		showUsage()
		os.Exit(0)
	}

	op := os.Args[1]
	switch op {
	case "-S", "--sync":
		checkRoot()
		if len(os.Args) < 3 {
			fmt.Println(styleErr.Render("✘ error:") + " package name required")
			os.Exit(1)
		}
		installPackage(os.Args[2:])

	case "-R", "--remove":
		checkRoot()
		if len(os.Args) < 3 {
			fmt.Println(styleErr.Render("✘ error:") + " package name required")
			os.Exit(1)
		}
		removePackage(os.Args[2:])

	case "-Ss", "--search":
		if len(os.Args) < 3 {
			fmt.Println(styleErr.Render("✘ error:") + " search query required")
			os.Exit(1)
		}
		searchPackage(os.Args[2:])

	case "-Syu", "--sync-upgrade":
		checkRoot()
		syncUpgrade()

	case "--help", "-h":
		showBanner()
		showUsage()

	case "--version", "-v":
		fmt.Printf("silvergrab v%s\n", version)

	default:
		fmt.Println(styleErr.Render("✘ error:") + fmt.Sprintf(" unknown option '%s'", op))
		fmt.Println(styleSubtl.Render("try: silvergrab --help"))
		os.Exit(1)
	}
}

func showBanner() {
	fmt.Println(styleLogo.Render(logo))
	fmt.Println(styleDim.Render("  SilverArch Package Manager  ·  v" + version))
	fmt.Println(styleDim.Render("  https://s1lverarch.github.io/linuxdistro/"))
	fmt.Println()
}

func showUsage() {
	usage := `Operations:
  ` + styleHeader.Render("-S, --sync") + `      install package(s)
  ` + styleHeader.Render("-R, --remove") + `    remove package(s)
  ` + styleHeader.Render("-Ss, --search") + `   search packages
  ` + styleHeader.Render("-Syu, --sync-upgrade") + ` sync and upgrade all

Examples:
  ` + styleDim.Render("sudo silvergrab -S firefox neovim") + `
  ` + styleDim.Render("sudo silvergrab -R vim") + `
  ` + styleDim.Render("silvergrab -Ss neovim") + `
  ` + styleDim.Render("sudo silvergrab -Syu") + `
`
	fmt.Println(usage)
}

func checkRoot() {
	if os.Geteuid() != 0 {
		fmt.Println(styleErr.Render("✘ error:") + " root required for this operation")
		fmt.Println(styleSubtl.Render("       try: sudo silvergrab " + os.Args[1]))
		os.Exit(1)
	}
}

func installPackage(pkgs []string) {
	fmt.Println()
	fmt.Println(styleHeader.Render("  Installing packages"))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))

	// show packages
	for i, pkg := range pkgs {
		fmt.Printf("  %s %s\n", styleNum.Render(fmt.Sprintf("[%d]", i+1)), pkg)
	}

	fmt.Println()
	fmt.Print(styleInfo.Render("  » continue? [Y/n] "))
	ans := readInput()
	if ans == "n" || ans == "N" {
		fmt.Println(styleSubtl.Render("  aborted"))
		return
	}

	fmt.Println()
	fmt.Println(styleHeader.Render("  Installing..."))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))

	cmd := exec.Command("pacman", append([]string{"-S"}, pkgs...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		fmt.Println(styleErr.Render("✘ error:") + " install failed")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))
	fmt.Println(styleOK.Render("  ✓ installation complete"))
	fmt.Println()
}

func removePackage(pkgs []string) {
	fmt.Println()
	fmt.Println(styleHeader.Render("  Removing packages"))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))

	for i, pkg := range pkgs {
		fmt.Printf("  %s %s\n", styleNum.Render(fmt.Sprintf("[%d]", i+1)), pkg)
	}

	fmt.Println()
	fmt.Print(styleInfo.Render("  » keep config files? [y/N] "))
	keepCfg := readInput()

	args := []string{"-R"}
	if keepCfg == "y" || keepCfg == "Y" {
		args = append(args, "-n")
		fmt.Println(styleSubtl.Render("  (keeping config files)"))
	}
	args = append(args, pkgs...)

	fmt.Println()
	fmt.Print(styleWarn.Render("  » proceed? [y/N] "))
	ans := readInput()
	if ans != "y" && ans != "Y" {
		fmt.Println(styleSubtl.Render("  aborted"))
		return
	}

	fmt.Println()
	fmt.Println(styleHeader.Render("  Removing..."))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))

	cmd := exec.Command("pacman", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		fmt.Println(styleErr.Render("✘ error:") + " remove failed")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))
	fmt.Println(styleOK.Render("  ✓ removal complete"))
	fmt.Println()
}

func searchPackage(query []string) {
	fmt.Println()
	fmt.Println(styleHeader.Render("  Searching packages"))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))
	fmt.Println()

	cmd := exec.Command("pacman", append([]string{"-Ss"}, query...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		fmt.Println(styleErr.Render("✘ error:") + " search failed")
		os.Exit(1)
	}
	fmt.Println()
}

func syncUpgrade() {
	fmt.Println()
	fmt.Println(styleHeader.Render("  System Upgrade"))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))
	fmt.Println(styleWarn.Render("  ⚠  this will upgrade all packages"))
	fmt.Println()
	fmt.Print(styleInfo.Render("  » proceed? [Y/n] "))
	ans := readInput()
	if ans == "n" || ans == "N" {
		fmt.Println(styleSubtl.Render("  aborted"))
		return
	}

	fmt.Println()
	fmt.Println(styleHeader.Render("  Syncing & Upgrading..."))
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))

	cmd := exec.Command("pacman", "-Syu")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println()
		fmt.Println(styleErr.Render("✘ error:") + " sync failed")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(styleDim.Render(strings.Repeat("─", 60)))
	fmt.Println(styleOK.Render("  ✓ system is up to date"))
	fmt.Println()
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
