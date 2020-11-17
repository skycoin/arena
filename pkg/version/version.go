package version

import "fmt"

var (
	// Release build args variable
	Release string
	// BuildTime build args variable
	BuildTime string
)

const unspecified = "unspecified"

// Print version
func Print() {
	if Release == "" {
		Release = unspecified
	}
	if BuildTime == "" {
		BuildTime = unspecified
	}
	fmt.Printf("Release: %s, BuildTime: %s \n", Release, BuildTime)
}
