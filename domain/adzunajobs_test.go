package domain

import (
	"sort"
	"testing"
	"time"
)

func TestSortMostToLeastRecentPosted(t *testing.T) {
	url := "www.khaluldidit.com"

	oldest := NewAdzunaJob(1, "Bayaz", "Magi", posted(2019, time.January), url, 0)

	secondOldest := NewAdzunaJob(2, "Khalul", "Magi", posted(2019, time.February), url, 0)

	secondNewest := NewAdzunaJob(3, "Yulwei", "Magi", posted(2020, time.January), url, 0)

	newest := NewAdzunaJob(4, "Yoru", "Magi", posted(2020, time.February), url, 0)

	jobs := AdzunaJobs([]AdzunaJob{oldest, newest, secondNewest, secondOldest})
	sort.Sort(jobs)

	if jobs[0] != newest {
		t.Fail()
	}

	if jobs[1] != secondNewest {
		t.Fail()
	}

	if jobs[2] != secondOldest {
		t.Fail()
	}

	if jobs[3] != oldest {
		t.Fail()
	}
}

func posted(year int, month time.Month) time.Time {
	return time.Date(year, month, 0,0,0,0,0, time.Local)
}