package data

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		AdmNm    string `json:"adm_nm"`
		AdmCd    string `json:"adm_cd"`
		AdmCd2   string `json:"adm_cd2"`
		Sgg      string `json:"sgg"`
		Sido     string `json:"sido"`
		SidoName string `json:"sidonm"`
		Temp     string `json:"temp"`
		Sggnm    string `json:"sggnm"`
		AdmCd8   string `json:"adm_cd8"`
	} `json:"properties"`
}

type HangJeongDong struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}
