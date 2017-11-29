# packer-post-processor-openstack-image-management
Packer post-processor plugin for OpenStack Image generation management

Inspired [packer-post-processor-amazon-ami-management](https://github.com/wata727/packer-post-processor-amazon-ami-management).

## Description
The packer OpenStack Image Management post-processor assists your image generation.
It delete and rotate old OpenStack Image.

## Installation
Download binary built for your architecture from [latest releases](https://github.com/nashiox/packer-post-processor-openstack-image-management/releases/latest).

For details on how to install the packer plugins, please read the following:
https://www.packer.io/docs/extending/plugins.html#installing-plugins

## Usage
The following example `sample.json`:

```json
{
  "builders": [{
    "type": "openstack",
    "identity_endpoint": "http://<destack-ip>:5000/v3",
    "tenant_name": "admin",
    "domain_name": "Default",
    "username": "admin",
    "password": "<your admin password>",
    "region": "RegionOne",
    "ssh_username": "root",
    "image_name": "sample-image",
    "source_image": "<image id>",
    "flavor": "m1.tiny",
    "insecure": "true"
  }],
  "provisioners":[{
    "type": "shell",
    "inline": [
      "echo 'running...'"
    ]
  }],
  "post-processors":[{
    "type": "openstack-image-management",
    "identity_endpoint": "http://<destack-ip>:5000/v3",
    "tenant_name": "admin",
    "domain_name": "Default",
    "username": "admin",
    "password": "<your admin password>",
    "region": "RegionOne",
    "insecure": "true",
    "identifier": "sample-image",
    "keep_releases": "2"
  }]
}
```

### configuration
Type: `openstack-image-management`

Required:
  - `identity_endpoint` (string)
    - The URL to the OpenStack Identity service. If not specified, Packer will use the environment variables OS_AUTH_URL, if set.
  - `username` (string)
    - The username or id used to connect to the OpenStack service. If not specified, Packer will use the environment variable OS_USERNAME, if set.
  - `password` (string)
    - The password used to connect to the OpenStack service. If not specified, Packer will use the environment variables OS_PASSWORD, if set.
  - `identifier` (string)
    - The Identifier of OpenStack Images.This plugin looks `image_name`.
  - `keep_releases` (interger)
    - The number of keep images.