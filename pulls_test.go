package gogit

import (
	"fmt"
	"testing"
)

func TestOpen(t *testing.T) {
	snowflake := Repo{"IoraHealth", "snowflake"}

	pulls, err := snowflake.OpenPulls()

	if err != nil {
		t.Errorf("Open failed due to error %v", err)
	}

	if len(pulls) <= 0 {
		t.Errorf("Open failed to fetch any open pulls")
	}

	for _, pull := range pulls {
		fmt.Printf("[number: %d, comments: %d, status: %s, octocatted: %v]", pull.State.Number, pull.State.CommentCount, pull.State.Status, pull.State.Octocatted)
	}
}
