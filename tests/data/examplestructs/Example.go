package examplestructs

type ProManagerMediaData struct {
	FieldData struct {
		CloudinaryPublicID string `json:"t_cloudinary_public_id"`
		DBRecidProduct     int    `json:"id_db_recid_product"`
		ImageNumber        int    `json:"c_imageNumber"`
	} `json:"fieldData"`
	Fid int `json:"recordId,string"`
}

type MediaResponse struct {
	Response struct {
		Data []ProManagerMediaData `json:"data"`
	} `json:"response"`
}
