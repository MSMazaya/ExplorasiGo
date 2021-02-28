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
		accountSid := "AC45d9bf7b4edeb91dbbd5ae152be0fe58"
		authToken := "7f1a7db524e4e3656d945545ce1a9463"
		urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
		fmt.Println(judge.Phone)
		v := url.Values{}
		v.Set("To", judge.Phone)
		v.Set("From", "+16152819135")
		v.Set("Body", "Tolong Mazaya! Dia mengalami kekerasan di Garut, Jl. Patriot No.31, Hubungi Polisi: 021-65303118 ")
		rb := *strings.NewReader(v.Encode())
		client := &http.Client{}

		req, _ := http.NewRequest("POST", urlStr, &rb)
		req.SetBasicAuth(accountSid, authToken)
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, _ := client.Do(req)
		fmt.Println(resp.Status)
	})

	r.Run()
}
