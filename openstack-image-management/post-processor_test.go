package openstackimagemanagement

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	th "github.com/gophercloud/gophercloud/testhelper"
	fakeclient "github.com/gophercloud/gophercloud/testhelper/client"
	"github.com/hashicorp/packer/packer"
)

func testUi() *packer.BasicUi {
	return &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	}
}

type imageEntry struct {
	ID   string
	JSON string
}

var imgs = []imageEntry{
	{
		ID: "packer-example-01",
		JSON: `{
				"status": "active",
				"name": "packer-example",
				"tags": [],
				"kernel_id": "e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
				"container_format": "ami",
				"created_at": "2015-07-15T11:43:35Z",
				"ramdisk_id": "8c64f48a-45a3-4eaa-adff-a8106b6c005b",
				"disk_format": "ami",
				"updated_at": "2015-07-15T11:43:35Z",
				"visibility": "public",
				"self": "/v2/images/07aa21a9-fa1a-430e-9a33-185be5982431",
				"min_disk": 0,
				"protected": false,
				"id": "07aa21a9-fa1a-430e-9a33-185be5982431",
				"size": 25165824,
				"file": "/v2/images/07aa21a9-fa1a-430e-9a33-185be5982431/file",
				"checksum": "eb9139e4942121f22bbc2afc0400b2a4",
				"owner": "cba624273b8344e59dd1fd18685183b0",
				"virtual_size": null,
				"min_ram": 0,
				"schema": "/v2/schemas/image",
				"hw_disk_bus": "scsi",
				"hw_disk_bus_model": "virtio-scsi",
				"hw_scsi_model": "virtio-scsi"
			}`,
	},
	{
		ID: "packer-example-02",
		JSON: `{
				"status": "active",
				"name": "packer-example",
				"tags": [],
				"container_format": "ari",
				"created_at": "2015-07-15T11:43:32Z",
				"size": 3740163,
				"disk_format": "ari",
				"updated_at": "2015-07-15T11:43:32Z",
				"visibility": "public",
				"self": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005b",
				"min_disk": 0,
				"protected": false,
				"id": "8c64f48a-45a3-4eaa-adff-a8106b6c005b",
				"file": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005b/file",
				"checksum": "be575a2b939972276ef675752936977f",
				"owner": "cba624273b8344e59dd1fd18685183b0",
				"virtual_size": null,
				"min_ram": 0,
				"schema": "/v2/schemas/image",
				"hw_disk_bus": "scsi",
				"hw_disk_bus_model": "virtio-scsi",
				"hw_scsi_model": "virtio-scsi"
			}`,
	},
	{
		ID: "packer-example-03",
		JSON: `{
				"status": "active",
				"name": "packer-example",
				"tags": [],
				"container_format": "aki",
				"created_at": "2015-07-15T11:43:29Z",
				"size": 4979632,
				"disk_format": "aki",
				"updated_at": "2015-07-15T11:43:30Z",
				"visibility": "public",
				"self": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
				"min_disk": 0,
				"protected": false,
				"id": "e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4",
				"file": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4/file",
				"checksum": "8a40c862b5735975d82605c1dd395796",
				"owner": "cba624273b8344e59dd1fd18685183b0",
				"virtual_size": null,
				"min_ram": 0,
				"schema": "/v2/schemas/image",
				"hw_disk_bus": "scsi",
				"hw_disk_bus_model": "virtio-scsi",
				"hw_scsi_model": "virtio-scsi"
			}`,
	},
	{
		ID: "cirros-0.3.4-x86_64-uec-kernel",
		JSON: `{
				"status": "active",
				"name": "cirros-0.3.4-x86_64-uec-kernel",
				"tags": [],
				"container_format": "aki",
				"created_at": "2015-07-15T11:43:29Z",
				"size": 4979632,
				"disk_format": "aki",
				"updated_at": "2015-07-15T11:43:30Z",
				"visibility": "public",
				"self": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba5",
				"min_disk": 0,
				"protected": false,
				"id": "e1b6edd4-bd9b-40ac-b010-8a6c16de4ba5",
				"file": "/v2/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba5/file",
				"checksum": "8a40c862b5735975d82605c1dd395796",
				"owner": "cba624273b8344e59dd1fd18685183b0",
				"virtual_size": null,
				"min_ram": 0,
				"schema": "/v2/schemas/image",
				"hw_disk_bus": "scsi",
				"hw_disk_bus_model": "virtio-scsi",
				"hw_scsi_model": "virtio-scsi"
			}`,
	},
	{
		ID: "cirros-0.3.4-x86_64-uec-ramdisk",
		JSON: `{
			"status": "active",
			"name": "cirros-0.3.4-x86_64-uec-ramdisk",
			"tags": [],
			"container_format": "ari",
			"created_at": "2015-07-15T11:43:32Z",
			"size": 3740163,
			"disk_format": "ari",
			"updated_at": "2015-07-15T11:43:32Z",
			"visibility": "public",
			"self": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005v",
			"min_disk": 0,
			"protected": false,
			"id": "8c64f48a-45a3-4eaa-adff-a8106b6c005v",
			"file": "/v2/images/8c64f48a-45a3-4eaa-adff-a8106b6c005v,/file",
			"checksum": "be575a2b939972276ef675752936977f",
			"owner": "cba624273b8344e59dd1fd18685183b0",
			"virtual_size": null,
			"min_ram": 0,
			"schema": "/v2/schemas/image",
			"hw_disk_bus": "scsi",
			"hw_disk_bus_model": "virtio-scsi",
			"hw_scsi_model": "virtio-scsi"
		}`,
	},
}

func TestPostProcessorEmptyImages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	ImageListHandler(t, []imageEntry{})

	p := OpenStackPostProcessor{conn: fakeclient.ServiceClient()}
	p.config.Identifier = "packer-example"
	p.config.KeepReleases = 3
	artifact := &packer.MockArtifact{}
	_, keep, err := p.PostProcess(testUi(), artifact)

	if !keep {
		t.Fatal("should keep")
	}

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestPostProcessorFewImages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	ImageListHandler(t, imgs[:2])

	p := OpenStackPostProcessor{conn: fakeclient.ServiceClient()}
	p.config.Identifier = "packer-example"
	p.config.KeepReleases = 3
	artifact := &packer.MockArtifact{}
	_, keep, err := p.PostProcess(testUi(), artifact)

	if !keep {
		t.Fatal("should keep")
	}

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestPostProcessorManyImages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	ImageListHandler(t, imgs)

	p := OpenStackPostProcessor{conn: fakeclient.ServiceClient()}
	p.config.Identifier = "packer-example"
	p.config.KeepReleases = 2
	artifact := &packer.MockArtifact{}
	_, keep, err := p.PostProcess(testUi(), artifact)
	if !keep {
		t.Fatal("should keep")
	}

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func ImageListHandler(t *testing.T, images []imageEntry) {
	th.Mux.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		limit := 10
		var err error
		if r.FormValue("limit") != "" {
			limit, err = strconv.Atoi(r.FormValue("limit"))
			if err != nil {
				t.Errorf("Error value for 'limit' parameter %v (error: %v)", r.FormValue("limit"), err)
			}

		}

		marker := ""
		newMarker := ""

		if r.Form["marker"] != nil {
			marker = r.Form["marker"][0]
		}

		t.Logf("limit = %v   marker = %v", limit, marker)

		selected := 0
		addNext := false
		var imageJSON []string

		fmt.Fprintf(w, `{"images": [`)

		for _, i := range images {
			if marker == "" || addNext {
				t.Logf("Adding image %v to page", i.ID)
				imageJSON = append(imageJSON, i.JSON)
				newMarker = i.ID
				selected++
			} else {
				if strings.Contains(i.JSON, marker) {
					addNext = true
				}
			}

			if selected == limit {
				break
			}
		}
		t.Logf("Writing out %v image(s)", len(imageJSON))
		fmt.Fprintf(w, strings.Join(imageJSON, ","))

		fmt.Fprintf(w, `],
			    "next": "/images?marker=%s&limit=%v",
			    "schema": "/schemas/images",
			    "first": "/images?limit=%v"}`, newMarker, limit, limit)

	})
	th.Mux.HandleFunc("/images/07aa21a9-fa1a-430e-9a33-185be5982431", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	th.Mux.HandleFunc("/images/8c64f48a-45a3-4eaa-adff-a8106b6c005b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	th.Mux.HandleFunc("/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba4", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	th.Mux.HandleFunc("/images/e1b6edd4-bd9b-40ac-b010-8a6c16de4ba5", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	th.Mux.HandleFunc("/images/8c64f48a-45a3-4eaa-adff-a8106b6c005v", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}
