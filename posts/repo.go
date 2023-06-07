package posts

type Repository interface {
	Save(interface{})
	Update(interface{})
	Delete(int)
	FindById(int, interface{}) (interface{}, error)
	FindAll() []interface{}
	FindByUserName(string) (interface{}, error)
}
