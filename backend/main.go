package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Judges struct {
	ID    string
	Phone string
}

func main() {
	dsn := "root:accessdenied@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	// Migrate the schema
	db.AutoMigrate(&Judges{})

	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		var judge Judges
		db.First(&judge, "id = ?", id)
		c.JSON(200, gin.H{
			"phone": judge.Phone,
		})
	})
	r.Run()

}
