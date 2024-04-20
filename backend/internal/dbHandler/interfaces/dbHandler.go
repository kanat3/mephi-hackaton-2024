package interfaces

type DbHandler interface {
	Create(object interface{})
	DeleteById(object interface{}, id string)
	SelectById(object interface{}, id string)

	Where(object interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	FindAll(object interface{})
}
