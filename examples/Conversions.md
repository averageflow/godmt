Original:

```go
const (
	NumberOne = 1
)

const (
	// Comment Test ExampleString
	ExampleString = "example"
)

const (
	//TrueConstant is true
	TrueConstant = true
)

type JustAType struct {
	Test string `json:"string"`
}

type NaughtyImport struct {
	DangerousField examplestructs.TestOne
	TestField      JustAType
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

// TestOne is a simple struct example
type TestOne struct {
    // CloudinaryPublicID This is a comment
    CloudinaryPublicID string `json:"t_cloudinary_public_id"`
    // This is a comment
    DBRecidProduct int `json:"id_db_recid_product"`
    // This is a comment
    ImageNumber int `json:"c_imageNumber"`
}

var (
    // NumberOne is a nice example
    // that spans multiple comment lines
    NumberOneVar = 1
)

var (
    // ExampleString is a nice example
    ExampleStringVar = "example"
    // This map also deserves a comment
    MapTestVar = map[string]string{
        "test":  "example",
        "test2": "example2",
        "test3": "example3",
    }
    MapTestIntVar = map[string]int{
        "test":  1,
        "test2": 2,
        "test3": 3,
    }
    MapTestInterfaceVar = map[string]interface{}{
        "test":  1,
        "test2": "",
        "test3": 1.234,
    }
)

var (
    //TrueConstant is true
    TrueVar = true
)

var (
    StringSlice = []string{
        "test",
        "test2",
    }
    NumberSlice = []int{
        1,
        2,
    }
    // InterfaceSlice is an interfaceSlice
    InterfaceSlice = []interface{}{
        "test",
        "test2",
    }
)
```

TypeScript:

```typescript
import { TestOne } from "examplestructs";


//TrueConstant is true
export const TrueVar: boolean = true;

export const StringSlice: string[] = [
	"test",
	"test2"
];

// InterfaceSlice is an interfaceSlice
export const InterfaceSlice: any[] = [
	"test",
	"test2"
];

export const NumberOne: number = 1;

// NumberOne is a nice example
// that spans multiple comment lines
export const NumberOneVar: number = 1;

// ExampleString is a nice example
export const ExampleStringVar: string = "example";

// This map also deserves a comment
export const MapTestVar: Record<string, string> = {
	"test": "example",
	"test2": "example2",
	"test3": "example3"
};

export const MapTestIntVar: Record<string, number> = {
	"test": 1,
	"test2": 2,
	"test3": 3
};

export const MapTestInterfaceVar: Record<string, any> = {
	"test": 1,
	"test2": "",
	"test3": 1.234
};

export const NumberSlice: number[] = [
	1,
	2
];

// Comment Test ExampleString
export const ExampleString: string = "example";

//TrueConstant is true
export const TrueConstant: boolean = true;


export interface JustAType {
	string: string;
}

export interface NaughtyImport {
	DangerousField: TestOne;
	TestField: JustAType;
}

export interface TestOne {
	// CloudinaryPublicID This is a comment
	t_cloudinary_public_id: string;
	// This is a comment
	id_db_recid_product: number;
	// This is a comment
	c_imageNumber: number;
}

export interface TestThree {
	test_field: string;
}

export interface TestFour extends TestThree, TestOne {
	another_field: boolean;
}

export interface TestFive {
	// This is a comment
	another_struct: TestOne;
	// This is a comment
	test_struct_easy: TestThree;
	// This is a comment
	another_field: boolean;
}

```

Swift:

```swift
// InterfaceSlice is an interfaceSlice
var InterfaceSlice: [Any] = [
	"test",
	"test2"
];

// Comment Test ExampleString
let ExampleString: String = "example"

//TrueConstant is true
let TrueConstant: Bool = true

// NumberOne is a nice example
// that spans multiple comment lines
let NumberOneVar: Int = 1

let MapTestIntVar: Dictionary<String, Int> = [
	"test": 1,
	"test2": 2,
	"test3": 3
]

let MapTestInterfaceVar: Dictionary<String, Any> = [
	"test": 1,
	"test2": "",
	"test3": 1.234
]

//TrueConstant is true
let TrueVar: Bool = true

var StringSlice: [String] = [
	"test",
	"test2"
];

let NumberOne: Int = 1

// ExampleString is a nice example
let ExampleStringVar: String = "example"

// This map also deserves a comment
let MapTestVar: Dictionary<String, String> = [
	"test": "example",
	"test2": "example2",
	"test3": "example3"
]

var NumberSlice: [Int] = [
	1,
	2
];


struct TestThree: Decodable {
	var test_field: String
}

struct TestFour: Decodable {
	var another_field: Bool
	var test_field: String
	// CloudinaryPublicID This is a comment
	var t_cloudinary_public_id: String
	// This is a comment
	var id_db_recid_product: Int
	// This is a comment
	var c_imageNumber: Int
}

struct TestFive: Decodable {
	// This is a comment
	var another_struct: TestOne
	// This is a comment
	var test_struct_easy: TestThree
	// This is a comment
	var another_field: Bool
}

struct JustAType: Decodable {
	var string: String
}

struct NaughtyImport: Decodable {
	var DangerousField: TestOne
	var TestField: JustAType
}

struct TestOne: Decodable {
	// CloudinaryPublicID This is a comment
	var t_cloudinary_public_id: String
	// This is a comment
	var id_db_recid_product: Int
	// This is a comment
	var c_imageNumber: Int
}

```