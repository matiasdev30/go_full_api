package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matiasdev30/go_api/db"
	"github.com/matiasdev30/go_api/models"
	"github.com/matiasdev30/go_api/utils"
)

// Create tasks
// @Summary Create task for works
// @Schemes
// @Description Endpoint for create task
// @Tags CreateTask
// @Accept json
// @Produce json
// @Param request body models.TaskUpdateRequest true "Request body"
// @Success 200 {object} models.TaskUpdateRequest
// @Router /task/createTask [post]
// @Security Bearer
func TaskHandler(ctx *gin.Context) {

	var data = models.Task{}

	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "cannot bind json "+err.Error())
		return
	}

	if err := db.Database.Create(&data).Error; err != nil {
		utils.SendErro(ctx, http.StatusInternalServerError, "error created user"+err.Error())
		return
	}

	utils.SendSucess(ctx, http.StatusOK, data, "data")

}

// List tasks
// @Summary List all tasks
// @Schemes
// @Description Endpoint for list task
// @Tags GetTasks
// @Accept json
// @Produce json
// @Success 200 {object} []models.TaskUpdateRequest{}
// @Router /task/getTasks [get]
// @Security Bearer
func GetTaskHanlder(ctx *gin.Context) {

	tasks := []models.Task{}

	if err := db.Database.Find(&tasks).Error; err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, err.Error())
		return
	}

	utils.SendSucess(ctx, http.StatusOK, tasks, "data")
}

// Delete tasks
// @Summary Delete task
// @Schemes
// @Description Endpoint for delete task
// @Tags DeleteTask
// @Accept json
// @Produce json
// @Param id query string true "Task identification"
// @Success 200 {object} models.TaskUpdateRequest{}
// @Router /task/deleteTask [delete]
// @Security Bearer
func DeleteTaskHandler(ctx *gin.Context) {

	id := ctx.Query("id")

	if id == "" {
		utils.SendErro(ctx, http.StatusBadRequest, "required query id")
		return
	}

	task := models.Task{}

	if err := db.Database.First(&task, id).Error; err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "task don't exist")
		return
	}

	if err := db.Database.Delete(&models.Task{}, id).Error; err != nil {
		utils.SendErro(ctx, http.StatusInternalServerError, "error delete task with id "+id)
		return
	}

	utils.SendSucess(ctx, http.StatusOK, task, "deleted-task")
}

// Update tasks
// @Summary Update task
// @Schemes
// @Description Endpoint for delete task
// @Tags UpdateTask
// @Accept json
// @Produce json
// @Param request body models.TaskUpdateRequest true "Request body"
// @Param id query string true "Task identification"
// @Success 200 {object} models.TaskUpdateRequest{}
// @Router /task/updateTask [put]
// @Security Bearer
func UpdateTaskHandler(ctx *gin.Context) {

	var requestTask models.TaskUpdateRequest

	err := ctx.ShouldBindJSON(&requestTask)

	if err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "cannot bind json "+err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		utils.SendErro(ctx, http.StatusBadRequest, "required query id")
		return
	}

	task := models.Task{}

	if err := db.Database.First(&task, id).Error; err != nil {
		utils.SendErro(ctx, http.StatusBadRequest, "task don't exist")
		return
	}

	if requestTask.Title != "" {
		task.Title = requestTask.Title
	}

	if requestTask.Description != "" {
		task.Description = requestTask.Description
	}

	if err := db.Database.Save(task).Error; err != nil {
		utils.SendErro(ctx, http.StatusInternalServerError, "error updating task with id "+id)
		return
	}

	if requestTask.Title == "" && requestTask.Description == "" {
		utils.SendErro(ctx, http.StatusInternalServerError, "insert data to update task")
		return
	}

	utils.SendSucess(ctx, http.StatusOK, task, "updated-task")
}
