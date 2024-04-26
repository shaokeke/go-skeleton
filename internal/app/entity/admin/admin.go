package admin

import "gorm.io/gen"

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int) (gen.T, error)

	// SELECT * FROM @@table WHERE username = @username
	GetByUserName(username string) (*gen.T, error)
}
