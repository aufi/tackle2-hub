package tag

import (
	"testing"

	"github.com/konveyor/tackle2-hub/api"
	"github.com/konveyor/tackle2-hub/test/api/client"
)

func TestTagCRUD(t *testing.T) {
	samples := Samples()

	for _, r := range samples {
		t.Run(r.Name, func(t *testing.T) {
			// Create.
			Should(t, Create, &r)

			// Get.
			got := api.Tag{}
			Should(t, Get, &got)
			if client.FlatEqual(got, &r) {
				t.Errorf("Different response error. Got %v, expected %v", got, r)
			}

			// Update.
			updated := api.Tag{
				Name:     "Updated " + r.Name,
				Category: r.Category,
			}
			Should(t, Update, &updated)

			Should(t, Get, &got)
			if got.Name != updated.Name {
				t.Errorf("Different response error. Got %s, expected %s", got.Name, updated.Name)
			}

			// Delete.
			Should(t, Delete, &r)

			// Ensure it is deleted
			//err := Client.Get(rPath, &r)
			//if err == nil {
			//	t.Errorf("Resource exits, but should be deleted: %v", r)
			//}
		})
	}
}

func TestTagList(t *testing.T) {
	samples := Samples()

	for i := range samples {
		Must(t, Create, &samples[i])
	}

	got := []api.Tag{}
	Should(t, List, &got)
	if client.FlatEqual(got, samples) {
		t.Errorf("Different response error. Got %v, expected %v", got, samples)
	}

	for _, r := range samples {
		Must(t, Delete, &r)
	}
}

func TestTagSeed(t *testing.T) {
	got := []api.Tag{}
	Should(t, List, &got)
	if len(got) < 1 {
		t.Errorf("Seed looks empty, but it shouldn't.")
	}
}
