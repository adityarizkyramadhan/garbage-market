package domain

import "gorm.io/gorm"

type MetaUser struct {
	gorm.Model
	Email           string           `json:"email"`
	Password        string           `json:"-"`
	Pin             string           `json:"pin"`
	User            User             `gorm:"foreignkey:IdMetaUser"`
	TokoSampah      TokoSampah       `gorm:"foreignkey:IdMetaUser"`
	PembayaranUsers []PembayaranUser `gorm:"foreignkey:IdMetaUser"`
	BasketJualans   []BasketJualan   `gorm:"foreignkey:IdMetaUser"`
}

type User struct {
	gorm.Model
	IdMetaUser   uint   `json:"id_meta_user"`
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	Gender       string `json:"gender"`
	LinkFoto     string `json:"link_foto"`
}

type TokoSampah struct {
	gorm.Model
	IdMetaUser    uint           `json:"id_meta_user"`
	NamaToko      string         `json:"nama_toko"`
	AlamatToko    string         `json:"alamat_toko"`
	BarangJualans []BarangJualan `gorm:"foreignkey:IdTokoSampah"`
}

type BarangJualan struct {
	gorm.Model
	IdTokoSampah int    `json:"id_toko_sampah"`
	LinkFoto     string `json:"link_foto"`
	NamaBarang   string `json:"nama_barang"`
	HargaBarang  int    `json:"harga_barang"`
	StokBarang   int    `json:"stok_barang"`
	TipeBarang   string `json:"tipe_barang"`
	Deskripsi    string `json:"deskripsi"`
}

type PembayaranUser struct {
	gorm.Model
	IdMetaUser       uint   `json:"id_meta_user"`
	IdBarangJualan   int    `json:"id_barang_jualan"`
	JumlahPembelian  int    `json:"jumlah_pembelian"`
	JumlahPembayaran int    `json:"jumlah_pembayaran"`
	StatusPembayaran bool   `json:"status_pembayaran" gorm:"default:false"`
	LinkFoto         string `json:"link_foto"`
}

type BasketJualan struct {
	gorm.Model
	IdMetaUser     uint `json:"id_meta_user"`
	IdBarangJualan int  `json:"id_barang_jualan"`
	JumlahBarang   int  `json:"jumlah_barang"`
}
