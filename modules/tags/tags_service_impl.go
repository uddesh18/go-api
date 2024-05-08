package tags

import (
	"go-api/helper"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository TagsRepository
	Validate       *validator.Validate
}

// Create implements TagsService.
func (t *TagsServiceImpl) Create(tags CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := Tags{
		Name:     tags.Name,
		LastName: tags.LastName,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService.
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService.
func (t *TagsServiceImpl) FindAll() []TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []TagsResponse
	for _, value := range result {
		tag := TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService.
func (t *TagsServiceImpl) FindById(tagsId int) TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService.
func (t *TagsServiceImpl) Update(tags UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func NewTagsServiceImpl(tagRepository TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}
