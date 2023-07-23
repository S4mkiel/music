package entity

type PdfFile struct {
	FileName string `json:"file_name" gorm:"type:varchar(255);not null" valid:"required"`
	FileData []byte `json:"file_data" gorm:"type:bytea;not null" valid:"required"`
	Base     `gorm:"embedded"`
}
