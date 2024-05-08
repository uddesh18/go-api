package tags

type TagsRepository interface {
	Save(tags Tags)
	Update(tags Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags Tags, err error)
	FindAll() []Tags
}
