package entity

type Todos struct {
	Id             uint   `json:"id"`
	Nama           string `json:"nama"`
	TanggalMulai   string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
	Deskripsi      string `json:"deskripsi"`
}
