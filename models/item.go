package models

type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"id"`
	Item_code   int    `gorm:"not null; type:int" json:"itemCode"`
	Description string `gorm:"not null; type:varchar(255)" json:"description"`
	Quantity    int    `gorm:"not null; type:int" json:"quantity"`
	Order_id    uint
}
