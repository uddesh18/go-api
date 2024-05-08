package tags

type TagsService interface {
	Create(tags CreateTagsRequest)
	Update(tags UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) TagsResponse
	FindAll() []TagsResponse
}
