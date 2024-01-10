package domain

type Department struct {
	Id               int
	Name             string
	ParentDepartment *ParentDepartment
	Supervisor       *Supervisor
}

type ParentDepartment struct {
	Id   int
	Name string
}

type Supervisor struct {
	Id   int
	Name string
}
