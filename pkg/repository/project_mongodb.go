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

// FindAll returns all projects
func (p *ProjectRepositoryMongoDB) FindAll() ([]collection.Project, error) {
	var projects []collection.Project

	c := p.db.MongoDB.C("projects")
	err := c.Find(bson.M{}).All(&projects)
	return projects, err
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

// FindByID returns Project from ID
func (p *ProjectRepositoryMongoDB) FindByID(id string) (collection.Project, error) {
	var project collection.Project

	c := p.db.MongoDB.C("projects")
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&project)

	return project, err
}

// Add a project
func (p *ProjectRepositoryMongoDB) Add(project *collection.Project) error {
	return p.db.MongoDB.C("projects").Insert(project)
}

// Update a project
func (p *ProjectRepositoryMongoDB) Update(project *collection.Project) error {
	return p.db.MongoDB.C("projects").UpdateId(project.ID, project)
}

// Delete a project
func (p *ProjectRepositoryMongoDB) Delete(id string) error {
	return p.db.MongoDB.C("projects").RemoveId(bson.ObjectIdHex(id))
}
