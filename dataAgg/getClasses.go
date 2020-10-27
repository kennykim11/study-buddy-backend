package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	term := "2201"

	res, err := http.Get("https://tigercenter.rit.edu/tigerCenterApp/tc/courseCatalog")

	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//fmt.Printf("%s\n", body)

	type class struct {
		CourseID      string `json:"courseId"`
		Subject       string `json:"subject"`
		CatalogNumber string `json:"catalogNumber"`
		Title         string `json:"CourseTitleLong"`
	}
	type collegeType struct {
		Departments []struct {
			Classes []class `json:"classes"`
		} `json:"departments"`
	}

	ingested := map[string]collegeType{}
	if err := json.Unmarshal([]byte(body), &ingested); err != nil {
		panic("err != nil:" + err.Error())
	}
	courses := make(map[string]class)
	for _, college := range ingested {
		for _, dep := range college.Departments {
			for _, course := range dep.Classes {
				course.CatalogNumber = strings.TrimSpace(course.CatalogNumber)
				courses[course.CourseID] = course
				//fmt.Printf("%#v\n", course)
			}
		}
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(courses); err != nil {
		log.Println(err)
	}

	ioutil.WriteFile("../classes.json", buf.Bytes(), 0644)

	type courseSections struct {
		SectionID     string `json:"classNumber"`
		SectionNumber string `json:"classSection"`
		Instructor    string `json:"instructorFullName"`
		Meetings      []struct {
			DayTimes string `json:"daytimes"`
			Location string `json:"locationShort"`
		} `json:"meetings"`
	}

	type searchResults struct {
		Sections []courseSections `json:"searchResults"`
	}

	sectionsMap := make(map[string][]courseSections)

	for courseID := range courses {
		postQuery := map[string]map[string]string{"searchParams": {"query": courseID, "term": term}}
		postJSON, _ := json.Marshal(postQuery)
		res, _ := http.Post("https://tigercenter.rit.edu/tigerCenterApp/tc/class-search", "application/json", bytes.NewBuffer(postJSON))

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}

		ingested := searchResults{}
		if err := json.Unmarshal([]byte(body), &ingested); err != nil {
			panic("err != nil:" + err.Error())
		}
		fmt.Printf("%s\n", courseID)
		sectionsMap[courseID] = ingested.Sections
	}

	buf = new(bytes.Buffer)
	enc = json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(sectionsMap); err != nil {
		log.Println(err)
	}

	ioutil.WriteFile("../sections.json", buf.Bytes(), 0644)
}
