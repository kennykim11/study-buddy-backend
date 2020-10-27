package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	classesJSONFile, _ := ioutil.ReadFile("../downloaded.json")
	classesJSONString := string(classesJSONFile)
	r.GET("/classes", func(c *gin.Context) {
		//user := c.DefaultQuery("user", "NOUSER")
		c.String(200, classesJSONString)
	})

	type courseSections struct {
		SectionID     string `json:"classNumber"`
		SectionNumber string `json:"classSection"`
		Instructor    string `json:"instructorFullName"`
		Meetings      []struct {
			DayTimes string `json:"daytimes"`
			Location string `json:"locationShort"`
		} `json:"meetings"`
	}
	fileText, _ := ioutil.ReadFile("../sections.json")
	sectionsMap := map[string][]courseSections{}
	if err := json.Unmarshal([]byte(fileText), &sectionsMap); err != nil {
		panic("err != nil:" + err.Error())
	}
	r.GET("/sections", func(c *gin.Context) {
		c.JSON(200, sectionsMap[c.Request.URL.Query()["classId"][0]])
	})

	r.Run()
}
