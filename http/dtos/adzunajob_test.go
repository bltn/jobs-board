package dtos_test

import (
	"html/template"
	"jobs-board/http/dtos"
	"testing"
	"time"
)

func TestFormatsPostedCorrectly(t *testing.T) {
	posted, _ := time.Parse(time.RFC3339, "2020-07-30T15:04:05Z")
	expectedFormat := template.HTML("Jul 30")

	adzunaDTO := dtos.NewAdzunaJobDTO(
		"Person Who Knows Adequate Amount of C++",
		"Widget Inc.",
		posted,
		"www.widget.com",
	)

	if adzunaDTO.Posted() != expectedFormat {
		t.Logf("Expected: %s. Actual: %s.", expectedFormat, adzunaDTO.Posted())
		t.Fail()
	}
}

