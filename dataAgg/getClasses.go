package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//term := "2201"

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
		CourseID      int    `json:"courseId,string"`
		Subject       string `json:"subject"`
		CatalogNumber string `json:"catalogNumber"`
		Title         string `json:"CourseTitleLong"`
		Sections      []int
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
	courses := make(map[int]class)
	for _, college := range ingested {
		for _, dep := range college.Departments {
			for _, course := range dep.Classes {
				course.CatalogNumber = strings.TrimSpace(course.CatalogNumber)
				courses[course.CourseID] = course
				//fmt.Printf("%#v\n", course)
			}
		}
	}

	type courseSections struct {
		SectionID     int    `json:"classNumber,string"`
		SectionNumber string `json:"classSection"`
		Instructor    string `json:"instructorFullName"`
		Meetings      []struct {
			DayTimes string `json:"daytimes"`
			Location string `json:"locationShort"`
		} `json:"meetings"`
	}

	type searchResults struct {
		Sections []*courseSections `json:"searchResults"`
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.Database("studybuddy").Collection("sections")
	fmt.Printf("%+v\n", collection)

	// for courseID := range courses {
	// 	postQuery := map[string]map[string]string{"searchParams": {"query": strconv.Itoa(courseID), "term": term}}
	// 	postJSON, _ := json.Marshal(postQuery)
	// 	res, _ := http.Post("https://tigercenter.rit.edu/tigerCenterApp/tc/class-search", "application/json", bytes.NewBuffer(postJSON))

	// 	body, readErr := ioutil.ReadAll(res.Body)
	// 	if readErr != nil {
	// 		log.Fatal(readErr)
	// 	}

	// 	ingested := searchResults{}
	// 	if err := json.Unmarshal([]byte(body), &ingested); err != nil {
	// 		panic("err != nil:" + err.Error())
	// 	}
	// 	fmt.Printf("%d\n", courseID)

	// 	//MONGO STUFF
	// 	courseObject := courses[courseID]

	// 	for _, section := range ingested.Sections {

	// 		courseObject.Sections = append(courseObject.Sections, section.SectionID)

	// 		var result courseSections
	// 		err := collection.FindOne(context.TODO(), bson.M{"sectionid": section.SectionID}).Decode(&result)
	// 		if err != nil { //If not in the database
	// 			_, err := collection.InsertOne(context.TODO(), section)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		} else {
	// 			_, err := collection.ReplaceOne(context.TODO(),
	// 				bson.D{{"sectionid", section.SectionID}},
	// 				section)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 		}
	// 	}
	// 	courses[courseID] = courseObject

	// }

	//Writing classes file
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(courses); err != nil {
		log.Println(err)
	}

	fmt.Println(buf)

	err = ioutil.WriteFile("classes.json", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
