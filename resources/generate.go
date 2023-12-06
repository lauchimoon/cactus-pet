//go:generate file2byteslice -package=images -input=guy.png -output=./images/guy.go -var=Cactus_png
package resources

import (
    // Lulz
    _ "github.com/hajimehoshi/file2byteslice"
)
