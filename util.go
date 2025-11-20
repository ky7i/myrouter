package myrouter

func validatePath(path string) {
	if string(path[0]) != "/" {
		panic("Path must have the prefix '/'.")
	}
}
