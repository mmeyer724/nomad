package catalog

import (
	"github.com/hashicorp/nomad/drivers/exec"
)

// This file is where all builtin plugins should be registered in the catalog.
// Plugins with build restrictions should be placed in the appropriate
// register_XXX.go file.
func init() {
	//	RegisterDeferredConfig(rawexec.PluginID, rawexec.PluginConfig, rawexec.PluginLoader)
	Register(exec.PluginID, exec.PluginConfig)
	//	Register(qemu.PluginID, qemu.PluginConfig)
	//	Register(java.PluginID, java.PluginConfig)
	//	RegisterDeferredConfig(docker.PluginID, docker.PluginConfig, docker.PluginLoader)
}
