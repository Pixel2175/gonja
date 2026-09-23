// Package methods provides built-in method implementations for template types.
package methods

import "github.com/pixel2175/gonja/v2/exec"

var All = exec.Methods{
	Bool:  boolMethods,
	Str:   strMethods,
	Int:   intMethods,
	Float: floatMethods,
	Dict:  dictMethods,
	List:  listMethods,
}
