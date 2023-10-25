package dto

type TodosResponseBody struct {
	Nama           string `json:"nama" validate:"required"`
	TanggalMulai   string `json:"tanggal_mulai" validate:"required,datetime,ltefield=tanggal_selesai"`
	TanggalSelesai string `json:"tanggal_selesai" validate:"required,datetime,gtefield=tanggal_mulai"`
	Deskripsi      string `json:"deskripsi" validate:"required"`
}
