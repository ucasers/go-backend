package models

type CipherPair struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name    string `gorm:"size:255;not null;" json:"name"`
	Pwd     string `gorm:"size:100;not null;" json:"pwd"`
	Key     string `gorm:"size:255;not null;" json:"key"`
	OwnerID uint32 `gorm:"not null" json:"owner_id"`                          // Foreign key field
	Owner   User   `gorm:"foreignKey:OwnerID;references:ID;onDelete:CASCADE"` // Relationship
}
