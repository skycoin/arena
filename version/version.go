package version

var (
	// Version is semver version of the component (substituted during build)
	Version = "latest"
	// GitCommit is substituted with git hash during the build
	GitCommit = "HEAD"
)
