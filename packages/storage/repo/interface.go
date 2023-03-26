package repo

type Student struct {
}

type StudentData interface {
	GetStudents() []Student
	GetStudentByID(id int) (*Student, error)
	AddStudent(student Student) error
	RemoveStudentByID(id int) error
	UpdateStudent(id int)
}

type School struct {
	students []Student
}
