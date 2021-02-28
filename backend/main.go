package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Judges struct {
	ID    string
	Phone string
}

func apiCall(number string) {
	accountSid := "AC45d9bf7b4edeb91dbbd5ae152be0fe58"
	authToken := "a2d456698b19b0e14a564a1d593e4ad5"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	v.Set("To", number)
	v.Set("From", "+16152819135")
	v.Set("Body", "Brooklyn's in the house!")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
}

func main() {
	dsn := "root:accessdenied@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	db.AutoMigrate(&Judges{})

	r.GET("/get/:id", func(c *gin.Context) {
		id := c.Param("id")
		var judge Judges
		db.First(&judge, "id = ?", id)
		apiCall(judge.Phone)
	})

	r.Run()
}
