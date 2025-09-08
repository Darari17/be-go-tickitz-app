package repositories

import (
	"context"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db: db}
}

func (ar *AuthRepo) Login(ctx context.Context, email string) (*models.User, error) {
	sql := `SELECT id, email, password, role FROM users WHERE email = $1 LIMIT 1`

	var user models.User
	err := ar.db.QueryRow(ctx, sql, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *AuthRepo) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	sql := `
		INSERT INTO users (email, password, role, created_at)
		VALUES ($1, $2, $3, NOW())
		RETURNING id, created_at, updated_at
	`
	err := ar.db.QueryRow(ctx, sql, user.Email, user.Password, user.Role).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ar *AuthRepo) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	sql := `
		INSERT INTO profile (user_id, firstname, lastname, phone_number)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, firstname, lastname, phone_number
	`
	err := ar.db.QueryRow(ctx, sql,
		profile.UserID,
		profile.FirstName,
		profile.LastName,
		profile.PhoneNumber,
	).Scan(&profile.UserID, &profile.FirstName, &profile.LastName, &profile.PhoneNumber)

	if err != nil {
		return nil, err
	}
	return profile, nil
}
