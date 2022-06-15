package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/segmentio/ksuid"
	"net/http"
	"telrobot/models"
	"telrobot/partner"
	"telrobot/pkg/e"
	"telrobot/repositories"
	"telrobot/util/common"
)

type UserController struct {
	Users  *repositories.UserRepository
	Client *partner.Client
}

func NewUserController(r *repositories.UserRepository) *UserController {
	client := partner.NewClient()
	return &UserController{Users: r, Client: client}
}

func (uc *UserController) Create(c *gin.Context) {
	var body models.CreateUserRequest
	err := c.BindJSON(&body)
	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	model := models.User{
		Model:                common.Model{},
		Name:                 body.Name,
		Email:                body.Email,
		Phone:                body.Phone,
		Password:             body.Password,
		PasswordConfirmation: body.PasswordConfirmation,
	}

	uuid, err := uc.Client.Create(model)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	model.Uuid = uuid

	res, err := uc.Users.Create(&model)
	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}
	c.JSON(http.StatusOK, e.Result(res))
}

func (uc *UserController) List(c *gin.Context) {
	var body models.ListUserRequest
	err := c.BindJSON(&body)
	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	name := ""
	email := ""

	//docs, total, err := uc.Users.List(name, email, body.Page, body.Count)
	docs, total, err := uc.Users.ListBySql(name, email, body.Count, body.Page - 1)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	c.Header("X-Total", fmt.Sprint(total))
	c.JSON(http.StatusOK, e.Result(docs))
}

func (uc *UserController) Get(c *gin.Context) {
	uuid := c.Param("uuid")
	//uuid := fmt.Sprintf("%s", ksuid.New().String())
	//fmt.Println("uuid:", uuid)

	doc, err := uc.Users.GetByUuid(uuid)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	} else if doc == nil || doc.ID == 0 {
		code := http.StatusNotFound
		c.JSON(code, e.Error(code))
		return
	}

	_, err = uc.Client.Get(doc.Uuid)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	c.JSON(http.StatusOK, e.Result(doc))
}

func (uc *UserController) Update(c *gin.Context) {
	var body models.UpdateUserRequest
	err := c.BindJSON(&body)
	if err != nil {
		code := http.StatusBadRequest
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	uuid := c.Param("uuid")

	user := map[string]interface{}{
		"Name": body.Name,
		"Email": body.Email,
		"Phone": body.Phone,
		"Password": body.Password,
		"PasswordConfirmation": body.PasswordConfirmation,
	}

	affected, err := uc.Users.UpdateByUuid(uuid,user)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	} else if affected <= 0 {
		code := http.StatusNotFound
		c.JSON(code, e.Error(code))
		return
	}

	err = uc.Client.Update(uuid, user)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	//c.String(http.StatusOK, "")
	c.JSON(http.StatusOK, e.Result(""))
}

func (uc *UserController) Remove(c *gin.Context) {
	uuid := c.Param("uuid")

	//_, err := strconv.Atoi(id)
	//if err != nil {
	//	code := http.StatusBadRequest
	//	c.JSON(code, e.ErrorText(code, errors.New("id 参数不正确").Error()))
	//	return
	//}

	doc, err := uc.Users.GetByUuid(uuid)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	} else if doc.ID == 0 {
		code := http.StatusNotFound
		c.JSON(code, e.Error(code))
		return
	}

	affected, err := uc.Users.RemoveByUuid(uuid)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	} else if affected <= 0 {
		code := http.StatusNotFound
		c.JSON(code, e.Error(code))
		return
	}

	err = uc.Client.Remove(uuid)
	if err != nil {
		code := http.StatusInternalServerError
		c.JSON(code, e.ErrorText(code, err.Error()))
		return
	}

	c.JSON(http.StatusOK, e.Result(""))
	//c.String(http.StatusOK, "")
}
