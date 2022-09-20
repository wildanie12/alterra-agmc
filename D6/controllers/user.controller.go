package controllers

import (
	"agmc_d6/models"
	"agmc_d6/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	us *services.UserService
}

func NewUser(us *services.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) Index(c echo.Context) error {
	users, err := uc.us.FindAll()
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusInternalServerError,
			"message": "error creating a user: " + err.Error(),
			"data":    nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success getting a list of users",
		"data":    users,
	})
}

func (uc *UserController) Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "userID is invalid",
			"data":    nil,
		})
	}

	user, err := uc.us.Find(id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    code,
			"message": "error finding a user: " + err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success getting a user",
		"data":    user,
	})
}

func (uc *UserController) Store(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := uc.us.Insert(user)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusInternalServerError,
			"message": "error creating a user: " + err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success creating a user",
		"data":    user,
	})
}

func (uc *UserController) Update(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "userID is invalid",
			"data":    nil,
		})
	}

	user, err = uc.us.Update(id, user)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    code,
			"message": "error updating a user: " + err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success updating a user",
		"data": map[string]interface{}{
			"id": id,
		},
	})
}

func (uc *UserController) Delete(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    http.StatusBadRequest,
			"message": "userID is invalid",
			"data":    nil,
		})
	}

	err = uc.us.Delete(id)
	if err != nil {
		code := http.StatusInternalServerError
		if err.Error() == "cannot get user detail: record not found" {
			code = http.StatusBadRequest
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "error",
			"code":    code,
			"message": "error deleting a user: " + err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "Success deleting a user",
		"data": map[string]interface{}{
			"id": id,
		},
	})
}
