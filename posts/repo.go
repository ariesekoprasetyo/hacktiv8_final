package posts

type Repository interface {
	Save(interface{})
	Update(interface{})
	Delete(int, interface{})
	FindById(int, interface{}) (interface{}, error)
	FindAll(interface{}) interface{}
	FindByUserName(string) (interface{}, error)
}
