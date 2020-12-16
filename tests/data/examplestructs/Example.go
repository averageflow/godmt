package examplestructs

// TestOne is a simple struct example
type TestOne struct {
	// CloudinaryPublicID This is a comment
	CloudinaryPublicID string `json:"t_cloudinary_public_id"`
	// This is a comment
	DBRecidProduct int `json:"id_db_recid_product"`
	// This is a comment
	ImageNumber int `json:"c_imageNumber"`
}

type TestTwo struct {
	// This is a comment
	Fid       int `json:"recordId,string"`
	FieldData struct {
		// This is a comment
		CloudinaryPublicID string `json:"t_cloudinary_public_id"`
		// This is a comment
		DBRecidProduct int `json:"id_db_recid_product"`
		// This is a comment
		ImageNumber int `json:"c_imageNumber"`
	} `json:"fieldData"`
}

type TestThree struct {
	TestField string `json:"test_field"`
}

// TestStructEasy is a simple struct example
type TestFour struct {
	// This is a comment
	TestThree
	// This is a comment
	TestOne
	AnotherField bool `json:"another_field"`
}

// TestStructEasy is a simple struct example
type TestFive struct {
	// This is a comment
	OneMoreStruct TestOne `json:"another_struct"`
	// This is a comment
	TwoMoreStructs TestThree `json:"test_struct_easy"`
	// This is a comment
	AnotherField bool `json:"another_field"`
}
