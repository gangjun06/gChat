package db

type UserInfo struct {
	Username string `gorm:"type:varchar(100);NOT NULL"`
	Avatar   string `gorm:"type:varchar(300);NOT NULL"`
}

type ServerInfo struct {
	ID      int    `gorm:"AUTO_INCREMENT;unique_index;NOT NULL"`
	Address string `gorm:"type:varchar(50);NOT NULL"`
	Name    string `gorm:"type:varchar(100);NOT NULL"`
}
