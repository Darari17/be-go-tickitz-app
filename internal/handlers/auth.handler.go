package handlers

import (
	"log"
	"net/http"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/Darari17/be-go-tickitz-app/pkg"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authRepo *repositories.AuthRepo
}

func NewAuthHandler(authRepo *repositories.AuthRepo) *AuthHandler {
	return &AuthHandler{authRepo: authRepo}
}

// Login godoc
// @Summary     Login User
// @Description Login dengan email dan password, JWT disini
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body models.LoginRequest true "Login Request"
// @Router      /auth/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var body models.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	user, err := ah.authRepo.Login(ctx, body.Email)
	if err != nil || user == nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	var hash pkg.HashConfig
	valid, err := hash.CompareHashAndPassword(body.Password, user.Password)
	if err != nil || !valid {
		log.Println(err.Error())
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	claim := pkg.NewJWTClaims(user.ID, string(user.Role))

	token, err := claim.GenToken()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id":    user.ID,
				"email": user.Email,
				"role":  user.Role,
			},
		},
	})
}

// Register godoc
// @Summary     Register User
// @Description Daftar User baru beserta profile
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body body models.RegisterRequest true "Register Request"
// @Router      /auth/register [post]
func (ah *AuthHandler) Register(ctx *gin.Context) {
	var body models.RegisterRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	var hash pkg.HashConfig
	hash.UseRecommended()
	hashed, err := hash.GenHash(body.Password)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash password"})
		return
	}

	role := body.Role
	if role == "" {
		role = "user"
	}

	user := models.User{
		Email:    body.Email,
		Password: hashed,
		Role:     models.Role(role),
	}

	newUser, err := ah.authRepo.RegisterUser(ctx, &user)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
		return
	}

	profile := models.Profile{
		UserID:      newUser.ID,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		PhoneNumber: body.PhoneNumber,
	}

	newProfile, err := ah.authRepo.CreateProfile(ctx, &profile)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create profile"})
		return
	}
	newUser.Profile = *newProfile

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Register Success",
		"data": gin.H{
			"id":        newUser.ID,
			"email":     newUser.Email,
			"role":      newUser.Role,
			"firstname": newUser.Profile.FirstName,
			"lastname":  newUser.Profile.LastName,
			"phone":     newUser.Profile.PhoneNumber,
		},
	})
}
