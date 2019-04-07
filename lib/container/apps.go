package container

var Catalog = map[string]Application{
	"sh": Application{
		"sh": Container{
			Name:    "sh",
			Image:   "busybox",
			Command: []string{"/bin/sh"},
		},
	},
}
