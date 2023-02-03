package main

import (
	win_bridge "github.com/containernetworking/plugins/plugins/main/windows/win-bridge"
	win_overlay "github.com/containernetworking/plugins/plugins/main/windows/win-overlay"
	"os"
	"path/filepath"

	"github.com/containernetworking/plugins/plugins/ipam/host-local"
	"github.com/containernetworking/plugins/plugins/meta/flannel"
	"github.com/docker/docker/pkg/reexec"
)

func main() {
	os.Args[0] = filepath.Base(os.Args[0])
	reexec.Register("host-local", hostlocal.Main)
	reexec.Register("bridge", win_bridge.Main)
	reexec.Register("overlay", win_overlay.Main)
	reexec.Register("flannel", flannel.Main)
	reexec.Init()
}
