package controllers

import (
	"agmc_d3/lib"
	"agmc_d3/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userRepo *lib.UserRepository
}

// NewAuth return instance of AuthController.
func NewAuth(repo *lib.UserRepository) *AuthController {
	return &AuthController{
		userRepo: repo,
	}
}

type Request struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (ac *AuthController) Login(c echo.Context) error {
	req := Request{}
	c.Bind(&req)

	user, err := ac.userRepo.FindByUsername(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusInternalServerError,
			"message": "your credentials is invalid",
			"data":    nil,
		})
	}

	token, _ := middlewares.CreateToken(user.Email)

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "success",
		"code": http.StatusOK,
		"message": "Success authenticating a user",
		"data": map[string]interface{} {
			"token": token,
			"user": user,
		},
	})
}
