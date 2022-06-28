package devices

import (
	"block_1/modbus"
)

var ch_1 = []string{"rashod_per_par", "rashod_pit_voda"}
var ch_2 = []string{"P", "T_nitk_A", "T_nitk_B"}
var ch_3 = []string{"T_pitvod_kotl", "T_uhodgaz_ventur1", "T_uhodgaz_ventur2", "T_mazut", "T_uhodgaz_ventur3", "T_uhodgaz_ventur4"}
var ch_4 = []string{"T_vozdh_pered_rvpa", "T_vozdh_pered_rvpb", "T_vozdh_pered_tvpslev", "T_vozdh_pered_tvpsprav"}
var ch_5 = []string{"T_vozdh_posl_rvpa", "T_vozdh_posl_rvpb", "T_vozdh_posl_tvpslev", "T_vozdh_posl_tvpsprav", "T_ekonomiz_slev", "T_ekonomiz_sprav",
	"T_dym_zarvp_slev", "T_dym_zarvp_sprav", "T_dym_zatvp_slev", "T_dym_zatvp_sprav"}
var ch_6 = []string{"o2_slev"}
var ch_7 = []string{"o2_sprav"}
var ch_8 = []string{"rash_gaz"}
var ch_9 = []string{"temp_ostrpar_posl_sk"}
var ch_10 = []string{"davl_perpar_do_sk"}
var ch_11 = []string{"rash_ostrpar_tg"}
var ch_12 = []string{"vakum"}
var ch_13 = []string{"T_vod_vhod_kond", "T_vod_vih_kond", "T_par_ou1", "T_par_ou2", "T_kond_kn"}
var ch_14 = []string{"T_par_psg1_1", "T_par_psg1_2", "T_par_psg2", "T_voda_vhod_psg1", "T_voda_vihod_psg1", "T_voda_posle_psg12"}
var ch_15 = []string{"P_para_psg1"}
var ch_16 = []string{"P_para_psg2"}

type Device struct {
	Id          int
	Qan_chanels int32
	Request     []byte
	Chanels     []string
}

type F1772 struct {
	D Device
}

type Trm138 struct {
	D Device
}

type Tm5104d struct {
	D Device
}

type Irt5920n struct {
	D Device
}

var Dev_1 = F1772{Device{1, 2, []byte{0x01, 0x04, 0x02, 0x00, 0x00, 0x04, 0xf0, 0x71}, ch_1}}
var Dev_2 = F1772{Device{2, 3, []byte{0x02, 0x04, 0x02, 0x00, 0x00, 0x06, 0x71, 0x83}, ch_2}}
var Dev_3 = Trm138{Device{24, 6, []byte{0x18, 0x04, 0x00, 0x00, 0x00, 0x1E, 0x72, 0x0b}, ch_3}}
var Dev_4 = Trm138{Device{32, 4, []byte{0x20, 0x04, 0x00, 0x00, 0x00, 0x14, 0xf6, 0xb4}, ch_4}}
var Dev_5 = Tm5104d{Device{5, 10, []byte{0x05, 0x03, 0x05, 0x00, 0x00, 0x14, 0x44, 0x8d}, ch_5}}
var Dev_6 = Irt5920n{Device{6, 1, []byte{0xff, 0x3a, 0x36, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x34, 0x33, 0x37, 0x32, 0x32, 0x0d}, ch_6}}
var Dev_7 = Irt5920n{Device{7, 1, []byte{0xff, 0x3a, 0x37, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x33, 0x31, 0x36, 0x39, 0x31, 0x0d}, ch_7}}
var Dev_8 = Irt5920n{Device{8, 1, []byte{0xff, 0x3a, 0x38, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x33, 0x33, 0x39, 0x39, 0x35, 0x0d}, ch_8}}
var Dev_9 = Irt5920n{Device{9, 1, []byte{0xff, 0x3a, 0x39, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x32, 0x31, 0x39, 0x36, 0x32, 0x0d}, ch_9}}
var Dev_10 = Irt5920n{Device{10, 1, []byte{0xff, 0x3a, 0x31, 0x30, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x35, 0x33, 0x36, 0x31, 0x0d}, ch_10}}
var Dev_11 = Irt5920n{Device{11, 1, []byte{0xff, 0x3a, 0x31, 0x31, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x35, 0x30, 0x36, 0x37, 0x32, 0x0d}, ch_11}}
var Dev_12 = F1772{Device{12, 1, []byte{0x0c, 0x04, 0x02, 0x00, 0x00, 0x02, 0x71, 0x6e}, ch_12}}
var Dev_13 = Trm138{Device{40, 5, []byte{0x28, 0x04, 0x00, 0x00, 0x00, 0x19, 0x36, 0x39}, ch_13}}
var Dev_14 = Trm138{Device{48, 6, []byte{0x30, 0x04, 0x00, 0x00, 0x00, 0x1e, 0x74, 0x23}, ch_14}}
var Dev_15 = Irt5920n{Device{15, 1, []byte{0xff, 0x3a, 0x31, 0x35, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x31, 0x36, 0x38, 0x38, 0x31, 0x0d}, ch_15}}
var Dev_16 = Irt5920n{Device{16, 1, []byte{0xff, 0x3a, 0x31, 0x36, 0x3b, 0x31, 0x3b, 0x30, 0x3b, 0x32, 0x39, 0x34, 0x32, 0x35, 0x0d}, ch_16}}

func (f F1772) Get_data() []float32 {
	var response = f.D.Request[3 : f.D.Qan_chanels*4+3]
	for i := 0; i < int(f.D.Qan_chanels); i++ {
		response = append(response[:], response[2:4]...)
		response = append(response[4:], response[0:2]...)
	}
	var temp = modbus.Rtu_temp(response)
	return temp

}

func (tm Tm5104d) Get_data() []float32 {
	var response = tm.D.Request[3 : tm.D.Qan_chanels*4+3]
	var temp = modbus.Rtu_temp(response)
	return temp

}

func (tr Trm138) Get_data() []float32 {
	var response = tr.D.Request[3 : tr.D.Qan_chanels*10+3]
	for i := 0; i < int(tr.D.Qan_chanels); i++ {
		response = append(response[:4*i], response[4*i+6:]...)
	}
	var temp = modbus.Rtu_temp(response)
	return temp
}

func (i Irt5920n) Get_data() []float32 {
	var response = i.D.Request[4:10]
	var temp = modbus.Ascii_temp(response)
	return temp
}

func (f F1772) Get_device() Device {
	var device = f.D
	return device
}

func (tm Tm5104d) Get_device() Device {
	var device = tm.D
	return device
}

func (tr Trm138) Get_device() Device {
	var device = tr.D
	return device
}

func (i Irt5920n) Get_device() Device {
	var device = i.D
	return device
}

type Geter interface {
	Get_data() []float32
	Get_device() Device
}

var Devices = [...]Geter{Dev_1, Dev_2, Dev_3, Dev_4, Dev_5, Dev_6, Dev_7, Dev_8, Dev_9, Dev_10, Dev_12, Dev_13, Dev_14, Dev_15, Dev_16}
