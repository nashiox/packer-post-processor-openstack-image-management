package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/packer/packer/plugin"
	openstackimagemanagement "github.com/nashiox/packer-post-processor-openstack-image-management/openstack-image-management"
)

var (
	Version  = "unset"
	Revision = "unset"
)

func main() {
	var version bool
	flag.BoolVar(&version, "version", false, "Show version.")
	flag.Parse()

	if version {
		fmt.Printf("Version: %s, Revision: %s", Version, Revision)
		os.Exit(0)
	}

	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	if err := server.RegisterPostProcessor(new(openstackimagemanagement.OpenStackPostProcessor)); err != nil {
		panic(err)
	}
	server.Serve()
}
