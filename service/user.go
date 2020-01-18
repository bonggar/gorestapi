package service

import (
	"github.com/bonggar/gorestapi/database"
	"github.com/bonggar/gorestapi/helper"
	"github.com/bonggar/gorestapi/model"

	"github.com/gin-gonic/gin"
)

//isUniqueEmail : check if the email address already taken
func isUniqueEmail(user model.User, isUpdate bool) bool {
	var res model.User
	db := database.GetDB()
	if isUpdate {
		db.Where("email = ? AND id <> ?", user.Email, user.ID).First(&res)
	} else {
		db.Where("email = ?", user.Email).First(&res)
	}

	if res.ID != 0 {

		return false
	}

	return true
}

//isUniquePhone : check if the phone number already taken
func isUniquePhone(user model.User, isUpdate bool) bool {
	var res model.User
	db := database.GetDB()
	if isUpdate {
		db.Where("phone = ? AND id <> ?", user.Phone, user.ID).First(&res)
	} else {
		db.Where("phone = ?", user.Phone).First(&res)
	}

	if res.ID != 0 {

		return false
	}

	return true
}

//GetUser : get a user which find by ID
func GetUser(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var user model.User
	db.First(&user, id)

	if user.ID != 0 {
		helper.RespondJSON(c, 200, "", user, nil)
	} else {
		helper.RespondJSON(c, 404, "User not found", nil, nil)
	}
}

//GetUsers : get all user
func GetUsers(c *gin.Context) {
	var users []model.User
	db := database.GetDB()
	db.Find(&users)

	helper.RespondJSON(c, 200, "", users, nil)
}

//CreateUser : create a user
func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBind(&user); err != nil {
		errorFields := helper.ConstructErrors(err)
		helper.RespondJSON(c, 422, "Could not create user", nil, errorFields)
		return
	}

	db := database.GetDB()
	//manually check email unique constraint
	if !isUniqueEmail(user, false) {
		errorFields := [1]helper.ErrorField{{ID: "email", Message: "Email already taken"}}
		helper.RespondJSON(c, 422, "Could not create user", nil, errorFields)
		return
	}

	//manually check phone unique constraint
	if !isUniquePhone(user, false) {
		errorFields := [1]helper.ErrorField{{ID: "phone", Message: "Phone already taken"}}
		helper.RespondJSON(c, 422, "Could not create user", nil, errorFields)
		return
	}

	//finally create the user
	db.Create(&user)

	if user.ID != 0 {
		helper.RespondJSON(c, 201, "User has been created successfully", user, nil)
	} else {
		helper.RespondJSON(c, 409, "Can not create user", nil, nil)
	}
}

//UpdateUser : edit a user which find by ID
func UpdateUser(c *gin.Context) {
	var user model.User
	db := database.GetDB()
	id := c.Params.ByName("id")
	db.First(&user, id)

	if user.ID != 0 {
		var updatedUser model.User
		c.ShouldBind(&updatedUser)

		//manually check email unique constraint
		if !isUniqueEmail(updatedUser, true) {
			errorFields := [1]helper.ErrorField{{ID: "email", Message: "Email already taken"}}
			helper.RespondJSON(c, 422, "Could not create user", nil, errorFields)
			return
		}

		//manually check phone unique constraint
		if !isUniquePhone(updatedUser, true) {
			errorFields := [1]helper.ErrorField{{ID: "phone", Message: "Phone already taken"}}
			helper.RespondJSON(c, 422, "Could not create user", nil, errorFields)
			return
		}

		//finally update the user
		db.Model(&user).Updates(updatedUser)

		helper.RespondJSON(c, 200, "User has been updated successfully", user, nil)
	} else {
		helper.RespondJSON(c, 404, "User not found", nil, nil)
	}
}

//DeleteUser : delete a user which find by ID
func DeleteUser(c *gin.Context) {
	var user model.User
	db := database.GetDB()
	id := c.Params.ByName("id")
	db.First(&user, id)

	if user.ID != 0 {
		db.Delete(&user)
		helper.RespondJSON(c, 200, "User has been deleted successfully", nil, nil)
	} else {
		helper.RespondJSON(c, 404, "User not found", nil, nil)
	}
}

//OptionsUser : supporting options for CORS
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
