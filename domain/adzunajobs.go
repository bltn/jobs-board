package domain

type AdzunaJobs []AdzunaJob

func (a AdzunaJobs) Len() int {
	return len(a)
}

func (a AdzunaJobs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AdzunaJobs) Less(i, j int) bool {
	return a[i].posted.After(a[j].posted)
}