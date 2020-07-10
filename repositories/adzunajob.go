package repositories

import (
	"context"
	"github.com/jackc/pgx/v4"
	"jobs-board/domain"
	"time"
)

type AdzunaJobRepository interface {
	FindAll() (domain.AdzunaJobs, error)
}

type adzunaJobRepository struct {
	dbConnection *pgx.Conn
}

func NewAdzunaJobRepository(driver *pgx.Conn) AdzunaJobRepository {
	return &adzunaJobRepository{driver}
}

func (r adzunaJobRepository) FindAll() (domain.AdzunaJobs, error) {
	rows, err := r.dbConnection.Query(context.Background(), "SELECT * FROM ADZUNA_JOBS;")

	if err != nil {
		return nil, err
	}

	var adzunaJobs []domain.AdzunaJob

	for rows.Next() {
		var id int
		var title string
		var company string
		var posted time.Time
		var url string
		var createdTsp time.Time
		var updatedTsp time.Time
		var adzunaId int

		err := rows.Scan(&id, &title, &company, &posted, &url, &createdTsp, &updatedTsp, &adzunaId)

		if err != nil {
			return nil, err
		}

		adzunaJob := domain.NewAdzunaJob(id, title, company, posted, url, adzunaId)
		adzunaJobs = append(adzunaJobs, adzunaJob)
	}

	return adzunaJobs, nil
}
