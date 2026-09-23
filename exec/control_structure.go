package exec

import (
	"github.com/pixel2175/gonja/v2/nodes"
)

type ControlStructure interface {
	nodes.ControlStructure
	Execute(*Renderer, *nodes.ControlStructureBlock) error
}
