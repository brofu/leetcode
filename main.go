package main

import (
	"flag"
	"fmt"

	"github.com/brofu/leetcode/version"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")

	//result := tree.Permute([]int{5, 4, 6, 2})
	//fmt.Println("Hello.", result)
}
