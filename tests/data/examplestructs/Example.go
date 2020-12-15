package examplestructs

type TestStructEasy struct {
	CloudinaryPublicID string `json:"t_cloudinary_public_id"`
	DBRecidProduct     int    `json:"id_db_recid_product"`
	ImageNumber        int    `json:"c_imageNumber"`
}

/*type TestStructComplicated struct{
	Fid int `json:"recordId,string"`
	FieldData struct {
		CloudinaryPublicID string `json:"t_cloudinary_public_id"`
		DBRecidProduct     int    `json:"id_db_recid_product"`
		ImageNumber        int    `json:"c_imageNumber"`
	} `json:"fieldData"`
}*/
