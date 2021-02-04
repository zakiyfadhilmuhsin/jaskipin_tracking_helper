package models

import (
	"github.com/jinzhu/gorm"
)

type Tracking struct {
	gorm.Model
	ID               uint   `json:"-"`
	IdOrder          string `gorm:"primaryKey"`
	NamaPengirim     string `json:"nama_pengirim"`
	AlamatPengirim   string `json:"alamat_pengirim"`
	TelpPengirim     string `json:"telp_pengirim"`
	NamaPenerima     string `json:"nama_penerima"`
	AlamatPenerima   string `json:"alamat_penerima"`
	NegaraPenerima   string `json:"negara_penerima"`
	KodePos          string `json:"kode_pos"`
	TelpPenerima     string `json:"telp_penerima"`
	NoId             string `json:"no_id"`
	MitraExpedisi    string `json:"mitra_expedisi"`
	AwbNo            string `json:"awb_no"`
	Berat            string `json:"berat"`
	JenisLayanan     string `json:"jenis_layanan"`
	Panjang          string `json:"panjang"`
	Lebar            string `json:"lebar"`
	Tinggi           string `json:"tinggi"`
	Total            string `json:"total"`
	Ongkir           string `json:"ongkir"`
	BayarJaskipin    string `json:"bayar_jaskipin"`
	MetodePembayaran string `json:"metode_pembayaran"`
	NamaBank         string `json:"nama_bank"`
	Status           string `json:"status"`
	TglOrder         string `json:"tgl_order"`
	JamOrder         string `json:"jam_order"`
	JumlahKodi       string `json:"jumlah_kodi"`
	Admin            string `json:"admin"`
	Pickup           string `json:"pickup"`
	NamaAgen         string `json:"nama_agen"`
	KetBank          string `json:"ket_bank"`
	TglManifest      string `json:"tgl_manifest"`
	ProvPengirim     string `json:"prov_pengirim"`
	KotaPengirim     string `json:"kota_pengirim"`
	ProvPenerima     string `json:"prov_penerima"`
	KotaPenerima     string `json:"kota_penerima"`
}
