package examplestructs

// TestStructEasy is a simple struct example
type TestStructEasy struct {
	CloudinaryPublicID string `json:"t_cloudinary_public_id"`
	DBRecidProduct     int    `json:"id_db_recid_product"`
	ImageNumber        int    `json:"c_imageNumber"`
}

type TestStructComplicated struct {
	Fid       int `json:"recordId,string"`
	FieldData struct {
		CloudinaryPublicID string `json:"t_cloudinary_public_id"`
		DBRecidProduct     int    `json:"id_db_recid_product"`
		ImageNumber        int    `json:"c_imageNumber"`
	} `json:"fieldData"`
}

type AnotherStruct struct {
	TestField string `json:"test_field"`
}

// TestStructEasy is a simple struct example
type TestExtendedStructEasy struct {
	AnotherStruct
	TestStructEasy
	AnotherField bool `json:"another_field"`
}

// TestStructEasy is a simple struct example
type TestExtendedStructEasyAgain struct {
	OneMoreStruct  AnotherStruct  `json:"another_struct"`
	TwoMoreStructs TestStructEasy `json:"test_struct_easy"`
	AnotherField   bool           `json:"another_field"`
}
