package application

import (
	"testing"

	"github.com/konveyor/tackle2-hub/api"
	c "github.com/konveyor/tackle2-hub/test/api/client"
)

func TestApplicationCreate(t *testing.T) {
	samples := Samples()
	// Create on array of Applications calls subtest
	for _, r := range samples {
		t.Run(r.Name, func(t *testing.T) {

			err := Client.Post(api.ApplicationsRoot, &r)
			if err != nil {
				t.Errorf(err.Error()) // Error for standard test failure or failed assertion
			}

			// The Get test not included here, but in get_test.go

			// Clean
			c.Must(t, Delete(&r))
		})
	}
}

func TestApplicationNotCreateDuplicates(t *testing.T) {
	r := Samples()[0]

	// Create sample.
	err := Client.Post(api.ApplicationsRoot, &r)
	if err != nil {
		t.Errorf("Create error: %v", err.Error())
	}

	// Prepare Application with duplicate Name.
	dup := &api.Application{
		Name: r.Name,
	}

	// Try create the duplicate.
	err = Client.Post(api.ApplicationsRoot, &dup)
	if err == nil {
		t.Errorf("Created duplicate application: %v", dup)

		// Clean the duplicate.
		c.Must(t, Delete(dup))
	}

	// Clean.
	c.Must(t, Delete(&r))
}

func TestApplicationNotCreateWithoutName(t *testing.T) {
	// Prepare Application without Name.
	r := &api.Application{
		Name: "",
	}

	// Try create the duplicate Application.
	err := Client.Post(api.ApplicationsRoot, &r)
	if err == nil {
		t.Errorf("Created empty application: %v", r)

		// Clean.
		c.Must(t, Delete(r))
	}
}
