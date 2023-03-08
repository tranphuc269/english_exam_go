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

func (role UserRole) ToString() string {
	return [...]string{"Admin", "Student", "Lecturer"}[role-1]
}
