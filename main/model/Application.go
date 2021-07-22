package model

type Application struct {
	appId int
	name string
	owner string
}

func (a Application) getName() string {
	return a.name
}

func (a Application) getId() int {
	return a.appId
}


