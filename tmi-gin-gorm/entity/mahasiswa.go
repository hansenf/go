package entity

import "time"

type Mahasiswa struct {
	ID            int32     `gorm:"primary_key:auto_increment" json:"-"`
	Name          string    `gorm:"type:varchar(255)" json:"-"`
	Password      string    `gorm:"type:varchar(255)" json:"-"`
	NomorHP       string    `gorm:"type:varchar(255)" json:"-"`
	Email         string    `gorm:"type:varchar(255)" json:"-"`
	UrlFoto       string    `gorm:"type:varchar(255)" json:"-"`
	NamaLengkap   string    `gorm:"type:varchar(255)" json:"-"`
	TanggalLahir  string    `gorm:"type:varchar(255)" json:"-"`
	JenisKelamin  string    `gorm:"type:varchar(255)" json:"-"`
	University    string    `gorm:"type:varchar(255)" json:"-"`
	Nim           string    `gorm:"type:varchar(255)" json:"-"`
	Jurusan       string    `gorm:"type:varchar(255)" json:"-"`
	TahunMasuk    string    `gorm:"type:varchar(255)" json:"-"`
	KotaKabupaten string    `gorm:"type:varchar(255)" json:"-"`
	IDUser        int32     `json:"id_user"`
	CreatedAt     time.Time `json:"created_at"`
	User          User      `gorm:"foreignkey:ID; constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role"`
}
