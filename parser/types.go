package parser

// Dependency represents a single package in the tree and its dependencies.
type Dependency struct {
	Name         string
	Version      string
	Dependencies map[string]Dependency
	Labels       map[string]string
}

type NodeMeta struct {
	NodeVersion     string
	LockfileVersion int
	PackageManager  string
}

// PkgTree holds metadata about the tree and the root node of the tree.
type PkgTree struct {
	Dependency
	Type                 string
	PackageFormatVersion string
	Meta                 NodeMeta
	HasDevDependencies   bool
	Cyclic               bool
	Size                 int
}
