package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hupf3/ios-backend/models"
)

// CreateCourse 创建一个课程
func CreateCourse(context *gin.Context) {
	var course models.Course
	if err := context.BindJSON(&course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "binding error",
		})
		return
	}
	if _, err := models.CreateCourse(course); err != nil {
		fmt.Println("error")
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "succeed"})
}

// DeleteCourseByID 通过 ID 删除一个课程
func DeleteCourseByID(context *gin.Context) {
	param := context.Param("courseID")
	courseID, _ := strconv.Atoi(param)

	err := models.DeleteCourseByID(courseID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器错误: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除课程成功"})
}

// GetCourseByID 通过 ID 获取一个课程
func GetCourseByID(context *gin.Context) {
	param := context.Param("courseID")
	courseID, _ := strconv.Atoi(param)

	data, err := models.GetCourseByID(courseID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "服务器错误: " + err.Error(), "data": data})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": 200, "msg": "获取课程成功", "data": data})
}

// UpdateCourseByID 通过 ID 更新一个课程
func UpdateCourseByID(context *gin.Context) {
	param := context.Param("courseID")
	courseID, _ := strconv.Atoi(param)

	var course models.Course
	if err := context.BindJSON(&course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "binding error",
		})
		return
	}
	course.CourseID = courseID
	if _, err := models.UpdateCourse(course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "succeed"})
}
