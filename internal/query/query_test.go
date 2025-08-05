package query

import "testing"

func TestClaimTasks(t *testing.T) {
	got := ClaimTasks(3)
	expected := `
		UPDATE backlite_tasks
		SET
			claimed_at = $1,
			attempts = attempts + 1
		WHERE id IN ($2,$3,$4)
	`

	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
