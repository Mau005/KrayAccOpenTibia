package models

type Towns struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Pos_x uint32 `gorm:"column:posx"`
	Pos_y uint32 `gorm:"column:posy"`
	Pos_z uint32 `gorm:"column:posz"`
}
