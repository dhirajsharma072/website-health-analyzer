package dao

import (
	"log"

	mgo "github.com/globalsign/mgo"

	. "github.com/dhirajsharma072/website-health-analyzer/src/models"
)

type WebsiteDAO struct {
	HostName     string
	DatabaseName string
}

var db *mgo.Database

const COLLECTION string = "websites"

// Establish a connection to database
func (m *WebsiteDAO) Connect() {
	session, err := mgo.Dial(m.HostName)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.DatabaseName)
}

// Find list of websites
func (m *WebsiteDAO) FindAll() ([]Website, error) {
	var websites []Website
	err := db.C(COLLECTION).Find(nil).Sort("updatedAt").All(&websites)
	return websites, err
}

// Insert a website into database
func (m *WebsiteDAO) Insert(website Website) error {
	error := db.C(COLLECTION).Insert(&website)
	return error
}

// Delete an existing website
func (m *WebsiteDAO) Delete(website Website) error {
	err := db.C(COLLECTION).Remove(&website)
	return err
}

// Update an existing website
func (m *WebsiteDAO) Update(website Website) error {
	err := db.C(COLLECTION).UpdateId(website.ID, &website)
	return err
}

func (m *WebsiteDAO) PatchSite(match map[string]interface{}, update map[string]interface{}) bool {
	err := db.C(COLLECTION).Update(match, update)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
