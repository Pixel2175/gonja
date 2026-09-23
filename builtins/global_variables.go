package builtins

import "github.com/pixel2175/gonja/v2/exec"

var GlobalVariables = exec.NewContext(map[string]any{
	"gonja": map[string]any{
		"version": "v0.0.0+trunk",
	},
})
