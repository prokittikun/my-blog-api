package models

import (
	"fmt"
	"my-blog-api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllBlogs(blog *[]Blog) (err error) {
	if err = config.DB.Find(blog).Error; err != nil {
		return err
	}
	return nil
}

func CreateBlog(blog *Blog) (err error) {
	//add create time
	blog.CreatedAt = time.Now().String()
	blog.UpdatedAt = time.Now().String()
	if err = config.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

func GetBlogByID(blog *Blog, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(blog).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBlog(blog *Blog, id string) (err error) {
	fmt.Println(blog)
	config.DB.Save(blog)
	return nil
}

func DeleteBlog(blog *Blog, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(blog)
	return nil
}
