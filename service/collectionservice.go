package service

import (
	"tcgcollectionmanager/dto"
	"tcgcollectionmanager/persistence/entity"
	"tcgcollectionmanager/persistence/repo"
)

type CollectionService struct {
	Repo *repo.CollectionRepo
}

func toCollectionDTO(collection *entity.Collection) dto.CollectionDTO {
	return dto.CollectionDTO{
		Name:        collection.Name,
		Description: collection.Description,
		ID:          collection.ID,
		CreatedAt:   collection.CreatedAt.Local().Format("dd:mm:yyyy"),
		UpdatedAt:   collection.CreatedAt.Local().Format("dd:mm:yyyy"),
	}
}

func (c *CollectionService) Create(data *dto.CreateCollectionDTO) (dto.CollectionDTO, error) {
	var created entity.Collection
	created.Name = data.Name
	created.Description = data.Description
	c.Repo.Create(&created)
	return toCollectionDTO(&created), nil
}

func (c *CollectionService) GetAll() []dto.CollectionDTO {
	res := c.Repo.All()
	out := make([]dto.CollectionDTO, len(*res))
	for _, e := range *res {
		data := toCollectionDTO(&e)
		out = append(out, data)
	}
	return out
}

func (c *CollectionService) Get(id *dto.CollectionIdDTO) (dto.CollectionDTO, error) {

}
