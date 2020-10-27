const express = require('express');
const router = express()
const fetch = require('node-fetch')
const cors = require('cors')

courseList = {}
sectionList = {}

var downloaded = require('./downloaded.json');
var searchParams = {"searchParams":{"query":"115341","term":"2201","filterAnd":null,"isAdvanced":false,"campus":null,"session":null,"courseAttributeOptions":[],"career":null,"college":null,"component":null,"creditsMax":null,"creditsMin":null,"precision":null,"instructionType":null,"instructor":null,"subject":null}}

for (var college in downloaded){
  college = downloaded[college]
  courseList[college.collegeName] = {}
  for (var department of college.departments){
    //console.log(downloaded[college][department].classes)
    for (var course of department.classes){
      courseList[college.collegeName][course.courseId] = `${course.subject}-${course.catalogNumber.trim()} ${course.courseTitleLong}`
      // searchParams.searchParams.query = course.courseId
      // fetch('https://tigercenter.rit.edu/tigerCenterApp/tc/class-search',
      // {headers: { 'Content-Type': 'application/json' }, method: 'POST', body: JSON.stringify(searchParams) })
      // .then(res => res.status == 200 ? res.json() : console.log(res))
      // .then(json => {sectionList[course.courseId] = json})
      // .catch(err => console.log(err));
    }
  }
  
}

//console.log(courseList)
console.log(sectionList)

router.use(cors()) //Cross-origin resource sharing


router.get('/classes', (req, res) => {
  res.json(courseList);
});
router.get('/sections', (req, res) => {
  res.send(JSON.stringify(courseList[req.query]));
});

router.listen(5000, () => {
  console.log("Hi")
})