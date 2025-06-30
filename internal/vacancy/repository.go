package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

func NewVacancyRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, createdat) VALUES (@email, @role, @company, @salary, @type, @location, @createdat)`
	args := pgx.NamedArgs{
		"email":     form.Email,
		"role":      form.Role,
		"company":   form.Company,
		"salary":    form.Salary,
		"type":      form.Type,
		"location":  form.Location,
		"createdat": time.Now(),
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("невозможно  создать вакансию: %w", err)
	}
	return nil
}

func (r *VacancyRepository) GetAll() ([]Vacancy, error) {
	query := "SELECT * from vacancies"
	rows, err := r.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	vacancies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Vacancy])
	if err != nil {
		return nil, err
	}
	return vacancies, nil
}
