package stakeholdergroup

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/client"
)

var (
	// Setup Hub API client
	Client = client.Client
)

//
// Set of valid resources for tests and reuse.
var Samples = []*api.StakeholderGroup{
	{
		Name:        "Mgmt",
		Description: "Management stakeholder group.",
	},
	{
		Name:        "Engineering",
		Description: "Engineering team.",
	},
}

//
// Creates a copy of Samples for a test.
func CloneSamples() (samples []*api.StakeholderGroup) {
	raw, err := json.Marshal(Samples)
	if err != nil {
		panic("ERROR cloning samples")
	}
	json.Unmarshal(raw, &samples)
	return
}

//
// Create.
func Create(t *testing.T, r *api.StakeholderGroup) {
	err := Client.Post(api.StakeholderGroupsRoot, &r)
	if err != nil {
		t.Fatalf("Create fatal error: %v", err.Error())
	}
}

//
// Delete.
func Delete(t *testing.T, r *api.StakeholderGroup) {
	err := Client.Delete(fmt.Sprintf("%s/%d", api.StakeholderGroupsRoot, r.ID))
	if err != nil {
		t.Fatalf("Delete fatal error: %v", err.Error())
	}
}