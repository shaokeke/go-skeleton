package article

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE title = @title
	GetByTitle(account string) (*gen.T, error)
}
