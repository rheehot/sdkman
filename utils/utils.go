package utils

import (
	"io"
	"os"
	"os/exec"
	"runtime"
)

func Platform() string {
	platform := runtime.GOOS
	if platform == "windows" {
		platform = "msys_nt-10.0"
	}
	return platform
}

func Pager(pages io.ReadCloser) {
	pager := os.Getenv("PAGER")
	p := Platform()

	if pager == "" {
		if p == "msys_nt-10.0" {
			pager = "more"
		} else {
			pager = "less"
		}
	}
	c1 := exec.Command(pager)
	c1.Stdin = pages
	c1.Stdout = os.Stdout
	_ = c1.Start()
	_ = c1.Wait()
	defer pages.Close()
}
