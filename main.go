package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
)



type Derp struct {
	gorm.Model
	State           string
	PullRequest     string
	GithubEventJson string `sql:"type:json"`
}

func main() {
	var event1, event2 GithubEvent
	json.Unmarshal(readFile("resrources/1.json"), &event1)
	json.Unmarshal(readFile("resrources/2.json"), &event2)
	//
	login := "root:derp@tcp(localhost:3306)/prManager?&parseTime=true"

	db, err := gorm.Open("mysql", login)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//// Migrate the schema
	db.AutoMigrate(&Derp{})

	d1 := makeDBRow(event1)
	d2 := makeDBRow(event2)

	// Create
	//db.Create(d1)
	//db.Create(d2)

	//Read
	var r1,r2 Derp
	//There are many ways to format SQL queries, http://gorm.io/docs/query.html
	db.First(&r1, 1)
	//db.Find(&r2, "state = ?", "Opened")
	fmt.Printf("%+v\n" ,r1.ID)


	db.Model(&r1).Update("state", "Raghu was here")

	// Delete - delete product
	//db.Delete(&product)

	//Dumping this here so go dosnt bitch at me about unused vars
	fmt.Sprintf(d1.State,d2.State,r2.State)
}

func readFile(path string) []byte {
	dat, _ := ioutil.ReadFile(path)
	return dat
}

func makeDBRow (event GithubEvent) *Derp {
	e := &Derp{}

	e.State = event.State
	e.PullRequest = event.PullRequest.Title
	e.GithubEventJson = func(event GithubEvent) string {
		e, _ := json.Marshal(event)
		eventJson := string(e)
		return eventJson
		}(event)
	return e
}