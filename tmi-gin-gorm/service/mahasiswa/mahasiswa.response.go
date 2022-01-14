package _mahasiswa

import (
	"time"
	"tmi-gin/entity"

	_user "tmi-gin/service/user"
)

type MahasiswaResponse struct {
	ID            int32              `json:"id"`
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
	IDUser        int32              `json:"id_user"`
	CreatedAt     time.Time          `json:"created_at"`
	User          _user.UserResponse `json:"user,omitempty"`
}

func NewMahasiswaResponse(m entity.Mahasiswa) MahasiswaResponse {
	return MahasiswaResponse{
		ID:       m.ID,
		Name:     m.Name,
		Password: m.Password,
		User:     _user.NewUserResponse(m.User),
	}
}

func NewMahasiswaArrayResponse(ms []entity.Mahasiswa) []MahasiswaResponse {
	mRes := []MahasiswaResponse{}
	for _, v := range ms {
		p := MahasiswaResponse{
			ID:       v.ID,
			Name:     v.Name,
			Password: v.Password,
			User:     _user.NewUserResponse(v.User),
		}
		mRes = append(mRes, p)
	}
	return mRes
}
