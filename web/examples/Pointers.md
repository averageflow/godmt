# Pointers

Below you can find some examples of how pointers get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
// ScanPointers represents the details of a scanned struct.
type ScanPointers struct {
    // Doc deserves a comment
    Doc *[]string `json:"doc" binding:"required" validation:"required"`
    // Name also deserves a comment
    Name         *string                   `json:"name" binding:"required" validation:"required"`
    Fields       *bool                     `json:"fields" binding:"required" validation:"required"`
    InternalType *int                      `xml:"internalType" binding:"required" validation:"required"`
    Test         *map[string]int           `xml:"test"`
    MapOfSlices  *map[string][]int         `json:"tricky"`
    Insane       *map[int][]map[string]int `json:"insane"`
    MoreTest     *[]map[string]int         `xml:"moreTest"`
}
```

TypeScript: <a name="ts"></a>

```ts
export interface ScanPointers {
    // Doc deserves a comment
    doc: string[] | null;
    // Name also deserves a comment
    name: string | null;
    fields: boolean | null;
    internalType: number | null;
    test: Record<string, number> | null;
    tricky: Record<string, number[]> | null;
    insane: Record<number, Record<string, number>[]> | null;
    moreTest: Record<string, number>[] | null;
}
```

Swift: <a name="swift"></a>

```swift
struct ScanPointers: Decodable {
	// Doc deserves a comment
	var doc: [String?]?
	// Name also deserves a comment
	var name: String?
	var fields: Bool?
	var internalType: Int?
	var test: Dictionary<String, Int>?
	var tricky: Dictionary<String, [Int]>?
	var insane: Dictionary<Int, [Dictionary<String, Int>]>?
	var moreTest: [Dictionary<String, Int>]?

	enum CodingKeys: String, CodingKey {
		case doc = "doc"
		case name = "name"
		case fields = "fields"
		case internalType = "internalType"
		case test = "test"
		case tricky = "tricky"
		case insane = "insane"
		case moreTest = "moreTest"
	}
}
```

PHP: <a name="php"></a>

```php
class ScanPointers {
	/**
	 * Doc deserves a comment
	 * @var ??string[] $doc
	 */
	public array $doc;

	/**
	 * Name also deserves a comment
	 */
	public ?string $name;

	public ?bool $fields;

	public ?int $internalType;

	/**
	 * @var array $test
	 */
	public array $test;

	/**
	 * @var array $tricky
	 */
	public array $tricky;

	/**
	 * @var array $insane
	 */
	public array $insane;

	/**
	 * @var ?array $moreTest
	 */
	public array $moreTest;


	public function __construct(array $data) {
		$this->doc = $data['doc'];
		$this->name = $data['name'];
		$this->fields = $data['fields'];
		$this->internalType = $data['internalType'];
		$this->test = $data['test'];
		$this->tricky = $data['tricky'];
		$this->insane = $data['insane'];
		$this->moreTest = $data['moreTest'];
	}
}
```

JSON: <a name="json"></a>

```json
{
  "enums": {},
  "types": {
    "ScanPointers": {
      "doc": null,
      "name": "ScanPointers",
      "fields": [
        {
          "name": "Doc",
          "kind": "[]string",
          "tag": "`json:\"doc\" binding:\"required\" validation:\"required\"`",
          "doc": [
            "// Doc deserves a comment"
          ],
          "importedEntity": null,
          "subFields": null,
          "internalType": 5,
          "isPointer": true
        },
        {
          "name": "Name",
          "kind": "string",
          "tag": "`json:\"name\" binding:\"required\" validation:\"required\"`",
          "doc": [
            "// Name also deserves a comment"
          ],
          "importedEntity": null,
          "subFields": null,
          "internalType": 2,
          "isPointer": true
        },
        {
          "name": "Fields",
          "kind": "bool",
          "tag": "`json:\"fields\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 2,
          "isPointer": true
        },
        {
          "name": "InternalType",
          "kind": "int",
          "tag": "`xml:\"internalType\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 2,
          "isPointer": true
        },
        {
          "name": "Test",
          "kind": "map[string]int",
          "tag": "`xml:\"test\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1,
          "isPointer": true
        },
        {
          "name": "MapOfSlices",
          "kind": "map[string][]int",
          "tag": "`json:\"tricky\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1,
          "isPointer": true
        },
        {
          "name": "Insane",
          "kind": "map[int][]map[string]int",
          "tag": "`json:\"insane\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1,
          "isPointer": true
        },
        {
          "name": "MoreTest",
          "kind": "[]map[string]int",
          "tag": "`xml:\"moreTest\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 5,
          "isPointer": true
        }
      ],
      "internalType": 4
    }
  }
}
```