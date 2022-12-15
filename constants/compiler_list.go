package constants

import "github.com/ahmadyogi543/go-exec/types"

var COMPILERS map[string]types.Compiler = map[string]types.Compiler{
	"py": {
		Name:       "Python",
		Extension:  "py",
		Executable: "python",
		Version:    "3.10.8",
	},
	"php": {
		Name:       "PHP",
		Extension:  "php",
		Executable: "php",
		Version:    "8.1.13",
	},
}
