package handlers

import (
	"html/template"
	"jobs-board/domain"
	"jobs-board/http/dtos"
	"jobs-board/repositories"
	"log"
	"net/http"
	"path/filepath"
	"sort"
)

type JobsHandler struct {
	adzunaRepo repositories.AdzunaJobRepository
	jobListingTemplatePath string
}

func NewJobsHandler(adzunaRepo repositories.AdzunaJobRepository, jobListingTemplatePath string) *JobsHandler {
	return &JobsHandler{adzunaRepo: adzunaRepo, jobListingTemplatePath: jobListingTemplatePath}
}

func (j JobsHandler) Index(w http.ResponseWriter, r *http.Request) {
	adzunaJobs, err := j.adzunaRepo.FindAll()

	if err != nil {
		log.Printf("Unable to load adzuna jobs from database: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	} else {
		adzunaJobDTOs := adzunaJobDTOsSortedByDate(adzunaJobs)

		path, _ := filepath.Abs(j.jobListingTemplatePath)
		t := template.Must(template.New("job-listings.gohtml").ParseFiles(path))
		err = t.Execute(w, adzunaJobDTOs)

		if err != nil {
			panic(err)
		}
	}
}

func adzunaJobDTOsSortedByDate(jobs domain.AdzunaJobs) []dtos.AdzunaJobDTO {
	var adzunaJobDTOs []dtos.AdzunaJobDTO

	sort.Sort(jobs)
	for _, j := range jobs {
		dto := dtos.NewAdzunaJobDTO(j.Title(), j.Company(), j.Posted(), j.Url())
		adzunaJobDTOs = append(adzunaJobDTOs, dto)
	}

	return adzunaJobDTOs
}