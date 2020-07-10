package domain

import "time"

type AdzunaJob struct {
	id      int
	title   string
	company string
	posted time.Time
	url     string
	adzunaId int
}

func NewAdzunaJob(id int, title string, company string, posted time.Time, url string, adzunaId int) AdzunaJob {
	return AdzunaJob{id: id, title: title, company: company, posted: posted, url: url, adzunaId: adzunaId}
}

func (a AdzunaJob) Posted() time.Time {
	return a.posted
}

func (a AdzunaJob) Id() int {
	return a.id
}

func (a AdzunaJob) Title() string {
	return a.title
}

func (a AdzunaJob) Company() string {
	return a.company
}

func (a AdzunaJob) Url() string {
	return a.url
}