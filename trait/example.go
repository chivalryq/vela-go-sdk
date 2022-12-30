package trait

type expose struct {
	properties exposeSpec
}

type exposeSpec struct {
	port       int
	exposeType string
}
