package dtos

import (
	"html/template"
	"time"
)

type AdzunaJobDTO struct {
	title   template.HTML
	company template.HTML
	posted template.HTML
	url     template.URL
}

func NewAdzunaJobDTO(title string, company string, posted time.Time, url string) AdzunaJobDTO {
	return AdzunaJobDTO{
		title:   template.HTML(title),
		company: template.HTML(company),
		posted: template.HTML(posted.Format("Jan 02")),
		url:     template.URL(url),
	}
}

func (a AdzunaJobDTO) Posted() template.HTML {
	return a.posted
}

func (a AdzunaJobDTO) Title() template.HTML {
	return a.title
}

func (a AdzunaJobDTO) Company() template.HTML {
	return a.company
}

func (a AdzunaJobDTO) Url() template.URL {
	return a.url
}


