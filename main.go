package main

import (
	"github.com/hashicorp/packer/packer/plugin"
	openstackimagemanagement "github.com/nashiox/packer-post-processor-openstack-image-management/openstack-image-management"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterPostProcessor(new(openstackimagemanagement.OpenStackPostProcessor))
	server.Serve()
}
