package category

type Category struct {
	ID          uint64
	Name        string `gorm:"index"`
	Description string
	PostCount   int
}
