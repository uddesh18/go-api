package tags

import "gorm.io/gorm"

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	panic("unimplemented")
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() []Tags {
	panic("unimplemented")
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags Tags, err error) {
	panic("unimplemented")
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags Tags) {
	panic("unimplemented")
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags Tags) {
	panic("unimplemented")
}

func NewTagsREpositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
