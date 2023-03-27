package jobfunction

import (
	"testing"

	"github.com/konveyor/tackle2-hub/addon"
	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/client"
)

var (
	// Setup Hub API client
	Client = client.Client
)

//
// Set of valid resources for tests and reuse.
func Samples() (samples []api.JobFunction) {
	samples = []api.JobFunction{
		{
			Name: "Engineer",
		},
		{
			Name: "Manager",
		},
	}
	return
}

//
// Create a Tag.
func Create(t *testing.T, r *api.JobFunction) {
	err := Client.Post(api.JobFunctionsRoot, &r)
	if err != nil {
		t.Fatalf("Create fatal error: %v", err.Error())
	}
}

//
// Delete the Tag.
func Delete(t *testing.T, r *api.JobFunction) {
	err := Client.Delete(addon.Params{api.ID: r.ID}.Inject(api.JobFunctionRoot))
	if err != nil {
		t.Fatalf("Delete fatal error: %v", err.Error())
	}
}
