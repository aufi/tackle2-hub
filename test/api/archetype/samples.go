package archetype

import (
	"github.com/konveyor/tackle2-hub/api"
)

// Set of valid resources for tests and reuse.
var (
	MinimalArchetype = api.Archetype{
		Name:        "Minimal Archetype",
		Description: "Archetype minimal sample 1",
		Comments:    "Archetype comments",
	}
	Samples = []api.Archetype{MinimalArchetype}
)
