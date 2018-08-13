package repository

import (
	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/store"
	"gopkg.in/mgo.v2/bson"
)

// ProjectRepositoryMongoDB is project collection interface
type ProjectRepositoryMongoDB struct {
	db *store.DB
}

// NewProjectRepositoryMongoDB returns new object
func NewProjectRepositoryMongoDB(db *store.DB) *ProjectRepositoryMongoDB {
	return &ProjectRepositoryMongoDB{
		db: db,
	}
}

// FindProjectIDByAPIKey returns ProjectID from APIKey
func (p *ProjectRepositoryMongoDB) FindProjectIDByAPIKey(key string) (string, error) {
	var project collection.Project

	c := p.db.MongoDB.C("projects")
	err := c.Find(bson.M{"api_keys": key}).One(&project)

	return project.ID.Hex(), err
}

// FindProjectByID returns Project from ID
func (p *ProjectRepositoryMongoDB) FindProjectByID(id string) (collection.Project, error) {
	var project collection.Project

	c := p.db.MongoDB.C("projects")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&project)

	return project, err
}
