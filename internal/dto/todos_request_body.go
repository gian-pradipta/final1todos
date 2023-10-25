package dto

type TodosRequestBody struct {
	Nama           string `json:"nama" validate:"required"`
	TanggalMulai   string `json:"tanggal_mulai" validate:"required"`
	TanggalSelesai string `json:"tanggal_selesai" validate:"required"`
	Deskripsi      string `json:"deskripsi" validate:"required"`
}
