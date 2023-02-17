package handlers

import (
	"net/http"

	"github.com/SGDIEGO/ApplicationHA/internal/domain"
	"github.com/SGDIEGO/ApplicationHA/internal/ports"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService ports.UserServiceI
}

func UserHDL(userService ports.UserServiceI) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {

	users, err := uh.userService.GetAllUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(400, "Not users")
		return
	}

	ctx.JSON(200, users)

}

func (uh *UserHandler) GetUser(ctx *gin.Context) {

	id := ctx.Param("id") // Get the id from the url
	user, err := uh.userService.GetUser(id)

	if err != nil {
		ctx.AbortWithStatusJSON(400, "Not found user")
		return
	}

	ctx.JSON(200, user)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {

	var user domain.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	err := uh.userService.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, err.Error())
	}

	ctx.JSON(http.StatusCreated, user)
}

func (uh *UserHandler) EditUser(ctx *gin.Context) {
	var err error
	var user domain.User

	err = ctx.ShouldBind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Erro": err.Error(),
		})
	}

	id := ctx.Param("id") // Get the id from the url
	err = uh.userService.EditUser(id, user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(200, user)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var err error
	id := ctx.Param("id") // Get the id from the url
	err = uh.userService.DeleteUser(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.String(http.StatusAccepted, "User deleted succesfully")
}
