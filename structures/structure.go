package structures

type Employee struct {
	FirstName string
	LastName  string
	ID        int
}

func (e *Employee) FirstNameData() string {
	return e.FirstName
}
