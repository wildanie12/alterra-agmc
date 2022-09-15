package controllers

import (
	"agmc_d3/lib"

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
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Login(c echo.Context) error {
	req := Request{}
	c.Bind(&req)

	ac.userRepo.Find()
}
