package controllers

import (
	"my-blog-api/models"

	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context) {
	var blog []models.Blog
	err := models.GetAllBlogs(&blog)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, blog)
	}
}

func CreateBlog(c *gin.Context) {
	var blog models.Blog
	c.BindJSON(&blog)
	err := models.CreateBlog(&blog)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, blog)
	}
}

func GetBlogByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var blog models.Blog
	err := models.GetBlogByID(&blog, id)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, blog)
	}
}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	id := c.Params.ByName("id")
	err := models.GetBlogByID(&blog, id)
	if err != nil {
		c.JSON(404, blog)
	}
	c.BindJSON(&blog)
	err = models.UpdateBlog(&blog, id)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, blog)
	}
}

func DeleteBlog(c *gin.Context) {
	var blog models.Blog
	id := c.Params.ByName("id")
	err := models.DeleteBlog(&blog, id)
	if err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, blog)
	}
}
