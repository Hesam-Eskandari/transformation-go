package repository

type JSON interface {
	Marshal() (string, error)
	Unmarshal(str string) error
}

type RelationalModel interface {
	JSON
}
