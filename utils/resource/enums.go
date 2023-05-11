package resource

type UserRole int

const (
	Admin    UserRole = 1
	Student  UserRole = 2
	Lecturer UserRole = 3
)

func (role UserRole) IsValid() bool {
	switch role {
	case Admin, Student, Lecturer:
		return true
	}

	return false
}

func (role UserRole) ToInt() int {
	switch role {
	case Admin:
		return 1
	case Student:
		return 2
	case Lecturer:
		return 3
	}
	return 0
}

func (role UserRole) ToString() string {
	return [...]string{"Admin", "Student", "Lecturer"}[role-1]
}

type QuestionCase int

const (
	QuestionReading   QuestionCase = 1
	QuestionListening QuestionCase = 2
)
