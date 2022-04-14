package models

type DaftarAnak struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	UserId        uint   `json:"userid" form:"userid" gorm:"column:userid"`
	Gender        uint   `json:"gender" form:"gender" gorm:"column:gender"`
	NamaAnak      string `json:"namaanak" form:"namaanak" gorm:"column:namaanak"`
	NamaPanggilan string `json:"namapanggilan" form:"namapanggilan"`
	TanggalLahir  string `json:"tanggallahir" form:"tanggallahir"`
	PanjangBadan  uint   `json:"panjangbadan" form:"panjangbadan"`
	LingkarKepala uint   `json:"lingkarkepala" form:"lingkarkepala"`
	BeratBadan    uint   `json:"beratbadan" form:"beratbadan"`
}

func (DaftarAnak) TableName() string {
	return "anak"
}

type DaftarAnakKedua struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	UserId        uint   `json:"userid" form:"userid" gorm:"column:userid"`
	Gender        uint   `json:"gender" form:"gender" gorm:"column:genderkedua"`
	NamaAnak      string `json:"namaanak" form:"namaanak" gorm:"column:nama_anak_kedua"`
	NamaPanggilan string `json:"namapanggilan" form:"namapanggilan" gorm:"column:nama_panggilan_kedua"`
	TanggalLahir  string `json:"tanggallahir" form:"tanggallahir" gorm:"column:tanggal_lahir_kedua"`
	PanjangBadan  uint   `json:"panjangbadan" form:"panjangbadan" gorm:"column:panjang_badan_kedua"`
	LingkarKepala uint   `json:"lingkarkepala" form:"lingkarkepala" gorm:"column:lingkar_kepala_kedua"`
	BeratBadan    uint   `json:"beratbadankedua" form:"beratbadankedua" gorm:"column:berat_badan_kedua"`
}

func (DaftarAnakKedua) TableName() string {
	return "anakkedua"
}

type DaftarAnakKetiga struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	UserId        uint   `json:"userid" form:"userid" gorm:"column:userid"`
	Gender        uint   `json:"gender" form:"gender" gorm:"column:genderketiga"`
	NamaAnak      string `json:"namaanak" form:"namaanak" gorm:"column:nama_anak_ketiga"`
	NamaPanggilan string `json:"namapanggilan" form:"namapanggilan" gorm:"column:nama_panggilan_ketiga"`
	TanggalLahir  string `json:"tanggallahir" form:"tanggallahir" gorm:"column:tanggal_lahir_ketiga"`
	PanjangBadan  uint   `json:"panjangbadan" form:"panjangbadan" gorm:"column:panjang_badan_ketiga"`
	LingkarKepala uint   `json:"lingkarkepala" form:"lingkarkepala" gorm:"column:lingkar_kepala_ketiga"`
	BeratBadan    uint   `json:"beratbadan" form:"beratbadan" gorm:"column:berat_badan_ketiga"`
}

func (DaftarAnakKetiga) TableName() string {
	return "anakketiga"
}

type DaftarAnakSemua struct {
	Id                  uint   `json:"id" gorm:"column:id"`
	UserId              uint   `json:"userid" form:"userid" gorm:"column:userid"`
	Gender              uint   `json:"gender" form:"gender" gorm:"column:gender"`
	NamaAnak            string `json:"namaanak" form:"namaanak" gorm:"column:namaanak"`
	NamaPanggilan       string `json:"namapanggilan" form:"namapanggilan"`
	TanggalLahir        string `json:"tanggallahir" form:"tanggallahir"`
	PanjangBadan        uint   `json:"panjangbadan" form:"panjangbadan"`
	LingkarKepala       uint   `json:"lingkarkepala" form:"lingkarkepala"`
	BeratBadan          uint   `json:"beratbadan" form:"beratbadan"`
	GenderKedua         uint   `json:"gender[2]" form:"gender" gorm:"column:genderkedua"`
	NamaAnakKedua       string `json:"namaanak[2]" form:"namaanak" gorm:"column:nama_anak_kedua"`
	NamaPanggilanKedua  string `json:"namapanggilan[2]" form:"namapanggilan" gorm:"column:nama_panggilan_kedua"`
	TanggalLahirKedua   string `json:"tanggallahir[2]" form:"tanggallahir" gorm:"column:tanggal_lahir_kedua"`
	PanjangBadanKedua   uint   `json:"panjangbadan[2]" form:"panjangbadan" gorm:"column:panjang_badan_kedua"`
	LingkarKepalaKedua  uint   `json:"lingkarkepala[2]" form:"lingkarkepala" gorm:"column:lingkar_kepala_kedua"`
	BeratBadanKedua     uint   `json:"beratbadankedua[2]" form:"beratbadankedua" gorm:"column:berat_badan_kedua"`
	GenderKetiga        uint   `json:"gender[3]" form:"gender" gorm:"column:genderketiga"`
	NamaAnakKetiga      string `json:"namaanak[3]" form:"namaanak" gorm:"column:nama_anak_ketiga"`
	NamaPanggilanKetiga string `json:"namapanggilan[3]" form:"namapanggilan" gorm:"column:nama_panggilan_ketiga"`
	TanggalLahirKetiga  string `json:"tanggallahir[3]" form:"tanggallahir" gorm:"column:tanggal_lahir_ketiga"`
	PanjangBadanKetiga  uint   `json:"panjangbadan[3]" form:"panjangbadan" gorm:"column:panjang_badan_ketiga"`
	LingkarKepalaKetiga uint   `json:"lingkarkepala[3]" form:"lingkarkepala" gorm:"column:lingkar_kepala_ketiga"`
	BeratBadanKetiga    uint   `json:"beratbadan[3]" form:"beratbadan" gorm:"column:berat_badan_ketiga"`
}
