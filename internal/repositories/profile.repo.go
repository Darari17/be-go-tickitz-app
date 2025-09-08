package repositories

import (
	"context"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepo struct {
	db *pgxpool.Pool
}

func NewProfileRepo(db *pgxpool.Pool) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (pr *ProfileRepo) GetProfile(ctx context.Context, userID int) (*models.Profile, error) {
	var p models.Profile
	query := `
		SELECT user_id, firstname, lastname, phone_number
		FROM profile WHERE user_id = $1
	`
	err := pr.db.QueryRow(ctx, query, userID).
		Scan(&p.UserID, &p.FirstName, &p.LastName, &p.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (pr *ProfileRepo) UpdateProfile(ctx context.Context, profile models.Profile) error {
	query := `
		UPDATE profile
		SET firstname = $1, lastname = $2, phone_number = $3
		WHERE user_id = $4
	`
	_, err := pr.db.Exec(ctx, query,
		profile.FirstName,
		profile.LastName,
		profile.PhoneNumber,
		profile.UserID,
	)
	return err
}
