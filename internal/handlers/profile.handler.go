package handlers

import (
	"net/http"

	"github.com/Darari17/be-go-tickitz-app/internal/models"
	"github.com/Darari17/be-go-tickitz-app/internal/repositories"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileRepo *repositories.ProfileRepo
}

func NewProfileHandler(profileRepo *repositories.ProfileRepo) *ProfileHandler {
	return &ProfileHandler{profileRepo: profileRepo}
}

// GetProfile godoc
// @Summary     Get User Profile
// @Description Data profil user login
// @Tags        Profile
// @Produce     json
// @Success     200 {object} map[string]interface{} "Get Profile Success"
// @Failure     404 {object} map[string]interface{} "Profile not found"
// @Router      /profile [get]
func (ph *ProfileHandler) GetProfile(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	profile, err := ph.profileRepo.GetProfile(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": profile})
}

// UpdateProfile godoc
// @Summary     Update User Profile
// @Description Update data profil user yang sedang login
// @Tags        Profile
// @Accept      json
// @Produce     json
// @Param       body body models.Profile true "Profile data"
// @Success     200 {object} map[string]interface{} "Profile updated"
// @Failure     400 {object} map[string]interface{} "Invalid request body"
// @Failure     500 {object} map[string]interface{} "Failed to update profile"
// @Router      /profile [put]
func (ph *ProfileHandler) UpdateProfile(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")

	var profile models.Profile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	profile.UserID = userID

	if err := ph.profileRepo.UpdateProfile(ctx, profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update profile"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "profile updated"})
}
