package service

import (
	"fmt"
	"time"

	"github.com/neko-neko/goflippy/pkg/collection"
	"github.com/neko-neko/goflippy/pkg/repository"
)

type ProjectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (p *ProjectService) FetchAllProjects() ([]collection.Project, error) {
	projects, err := p.projectRepo.FindAll()
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return projects, err
}

func (p *ProjectService) FetchProject(id string) (collection.Project, error) {
	project, err := p.projectRepo.FindByID(id)
	if err != nil {
		err = NewStoreSystemError(err.Error())
	}

	return project, err
}

func (p *ProjectService) RegisterProject(project *collection.Project) error {
	if err := p.projectRepo.Add(project); err != nil {
		return NewStoreSystemError(err.Error())
	}

	return nil
}

func (p *ProjectService) UpdateProjectInfo(id string, name string, description string) (collection.Project, error) {
	project, err := p.projectRepo.FindByID(id)
	if err != nil {
		return project, NewResourceNotFoundError(fmt.Sprintf("project does not exists %s", id))
	}

	updated := false
	if len(name) != 0 {
		project.Name = name
		updated = true
	}
	if len(description) != 0 {
		project.Description = description
		updated = true
	}

	if updated {
		project.UpdatedAt = time.Now()

		err = p.projectRepo.Update(&project)
		if err != nil {
			err = NewStoreSystemError(err.Error())
		}
	}

	return project, err
}
