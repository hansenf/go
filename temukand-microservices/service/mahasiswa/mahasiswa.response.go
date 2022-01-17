package _mahasiswa

import (
	"time"
	"tmi-gin/entity"

	_user "tmi-gin/service/user"
)

type MahasiswaResponse struct {
	ID            int64              `json:"id"`
	Name          string             `json:"name"`
	Password      string             `json:"password"`
	NomorHP       string             `json:"email"`
	UrlFoto       string             `json:"url_foto"`
	NamaLengkap   string             `json:"nama_lengkap"`
	TanggalLahir  string             `json:"tanggal_lahir"`
	JenisKelamin  string             `json:"jenis_kelamin"`
	University    string             `json:"university"`
	Nim           string             `json:"nim"`
	Jurusan       string             `json:"jurusan"`
	TahunMasuk    string             `json:"tahun_masuk"`
	KotaKabupaten string             `json:"kota_kabupaten"`
	UserID        int64              `json:"id_user"`
	CreatedAt     time.Time          `json:"created_at"`
	User          _user.UserResponse `json:"user,omitempty"`
}

func NewMahasiswaResponse(m entity.Mahasiswa) MahasiswaResponse {
	return MahasiswaResponse{
		ID:            m.ID,
		Name:          m.Name,
		Password:      m.Password,
		NomorHP:       m.NomorHP,
		UrlFoto:       m.UrlFoto,
		NamaLengkap:   m.NamaLengkap,
		TanggalLahir:  m.TanggalLahir,
		JenisKelamin:  m.JenisKelamin,
		University:    m.University,
		Nim:           m.Nim,
		Jurusan:       m.Jurusan,
		TahunMasuk:    m.TahunMasuk,
		KotaKabupaten: m.KotaKabupaten,
		UserID:        m.UserID,
		CreatedAt:     m.CreatedAt,
		User:          _user.NewUserResponse(m.User),
	}
}

func NewMahasiswaArrayResponse(v []entity.Mahasiswa) []MahasiswaResponse {
	mRes := []MahasiswaResponse{}
	for _, v := range v {
		p := MahasiswaResponse{
			ID:            v.ID,
			Name:          v.Name,
			Password:      v.Password,
			NomorHP:       v.NomorHP,
			UrlFoto:       v.UrlFoto,
			NamaLengkap:   v.NamaLengkap,
			TanggalLahir:  v.TanggalLahir,
			JenisKelamin:  v.JenisKelamin,
			University:    v.University,
			Nim:           v.Nim,
			Jurusan:       v.Jurusan,
			TahunMasuk:    v.TahunMasuk,
			KotaKabupaten: v.KotaKabupaten,
			UserID:        v.UserID,
			CreatedAt:     v.CreatedAt,
			User:          _user.NewUserResponse(v.User),
		}
		mRes = append(mRes, p)
	}
	return mRes
}
