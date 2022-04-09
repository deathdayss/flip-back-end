package models

type Zone struct {
	ID   int    `gorm:"primaryKey;autoIncrement:false"` //this is GID
	Zone string `gorm:"primaryKey"`                     // GID + Zone serves as primary key
}
