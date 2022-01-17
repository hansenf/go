package dto

import "time"

type CreateMaidRequest struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	NomorHP       string `json:"nomor_hp"`
	Email         string `json:"email"`
	UrlFoto       string `json:"url_foto"`
	NamaLengkap   string `json:"nama_lengkap"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	University    string `json:"university"`
	Nim           string `json:"nim"`
	Jurusan       string `json:"jurusan"`
	TahunMasuk    string `json:"tahun_masuk"`
	KotaKabupaten string `json:"kota_kabupaten"`
}

type UpdateMaidRequest struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Password      string    `json:"password"`
	NomorHP       string    `json:"nomor_hp"`
	Email         string    `json:"email"`
	UrlFoto       string    `json:"url_foto"`
	NamaLengkap   string    `json:"nama_lengkap"`
	TanggalLahir  string    `json:"tanggal_lahir"`
	JenisKelamin  string    `json:"jenis_kelamin"`
	University    string    `json:"university"`
	Nim           string    `json:"nim"`
	Jurusan       string    `json:"jurusan"`
	TahunMasuk    string    `json:"tahun_masuk"`
	KotaKabupaten string    `json:"kota_kabupaten"`
	IDUser        int64     `json:"id_user"`
	CreatedAt     time.Time `json:"created_at"`
}
