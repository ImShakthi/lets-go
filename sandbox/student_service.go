package sandbox

type Student struct {
	Name string
	Age  int
}

type StudentService interface {
	CreateStudent(name string, age int) Student
}

type studentService struct{}

func NewStudentService() StudentService {
	return &studentService{}
}

func (s studentService) CreateStudent(name string, age int) Student {
	return Student{
		Name: name,
		Age:  age,
	}
}
