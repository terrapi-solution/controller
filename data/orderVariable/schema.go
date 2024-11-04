package orderVariable

type OrderVariable struct {
	Order int    `gorm:"not null"`
	Name  string `gorm:"not null"`
	Value string `gorm:"not null"`
}
