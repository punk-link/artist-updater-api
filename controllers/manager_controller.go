package controllers

import (
	"main/models/organizations"
	requests "main/requests/organizations"
	service "main/services/organizations"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddManager(ctx *gin.Context) {
	currentManager, err := getCurrentManager()
	if err != nil {
		Forbidden(ctx, err)
		return
	}

	var manager organizations.Manager
	if err := ctx.ShouldBindJSON(&manager); err != nil {
		UnprocessableEntity(ctx, err)
		return
	}

	result, err := service.AddManager(currentManager, manager)
	OkOrBadRequest(ctx, result, err)
}

func AddMasterManager(ctx *gin.Context) {
	var request requests.AddMasterManagerRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		UnprocessableEntity(ctx, err)
		return
	}

	result, err := service.AddMasterManager(request)
	OkOrBadRequest(ctx, result, err)
}

func GetManager(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	result, err := service.GetManager(id)
	OkOrBadRequest(ctx, result, err)
}

func GetManagers(ctx *gin.Context) {
	currentManager, err := getCurrentManager()
	if err != nil {
		Forbidden(ctx, err)
		return
	}

	result := service.GetOrganizationManagers(currentManager.OrganizationId)
	Ok(ctx, result)
}

func ModifyManager(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		BadRequest(ctx, err.Error())
		return
	}

	var manager organizations.Manager
	if err := ctx.ShouldBindJSON(&manager); err != nil {
		UnprocessableEntity(ctx, err)
		return
	}

	result, err := service.ModifyManager(manager, id)
	OkOrBadRequest(ctx, result, err)
}
