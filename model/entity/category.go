package entity

type Category struct {
	Id          int64  `pg:"id, pk"`
	Name        string `pg:"name"`
	Description string `pg:"description"`
}
