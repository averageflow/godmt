# Structs

Below you can find some examples of how structs get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
// ScannedType represents a basic entity to be translated.
// More specifically const and var items.
type ScannedType struct {
    Name         string      `json:"name"`
    Kind         string      `json:"kind"`
    Value        interface{} `json:"value"`
    Doc          []string    `json:"doc"`
    InternalType int         `json:"internalType"`
}

type ExtendedScannedType struct {
    ScannedType
    AnExtraField bool `json:"anExtraField"`
}

// ScannedStruct represents the details of a scanned struct.
type ScannedStruct struct {
    Doc          []string `json:"doc" binding:"required" validation:"required"`
    Name         string   `json:"name" binding:"required" validation:"required"`
    Fields       []bool   `json:"fields" binding:"required" validation:"required"`
    InternalType int      `xml:"internalType" binding:"required" validation:"required"`
}

type EmbeddedType struct {
    ID   string `json:"id"`
    Data struct {
        Test string `json:"test"`
    } `json:"data"`
}
```

TypeScript: <a name="ts"></a>

```ts
export interface EmbeddedType {
    id: string;
    data: {
        test: string;
    }
}

export interface ExtendedScannedType extends ScannedType {
    anExtraField: boolean;
}

export interface ScannedStruct {
    doc: string[];
    name: string;
    fields: boolean[];
    internalType: number;
}

export interface ScannedType {
    name: string;
    kind: string;
    doc: string[];
    internalType: number;
}
```

Swift: <a name="swift"></a>

```swift
struct EmbeddedType: Decodable {
	var id: String
	struct data {
		var test: String;
	}

	enum CodingKeys: String, CodingKey {
		case id = "id"
		case data = "data"
	}
}

struct ExtendedScannedType: Decodable {
	var anExtraField: Bool
	var name: String
	var kind: String
	var doc: [String]
	var internalType: Int

	enum CodingKeys: String, CodingKey {
		case anExtraField = "anExtraField"
		case name = "name"
		case kind = "kind"
		case doc = "doc"
		case internalType = "internalType"
	}
}

struct ScannedStruct: Decodable {
	var doc: [String]
	var name: String
	var fields: [Bool]
	var internalType: Int

	enum CodingKeys: String, CodingKey {
		case doc = "doc"
		case name = "name"
		case fields = "fields"
		case internalType = "internalType"
	}
}

struct ScannedType: Decodable {
	var name: String
	var kind: String
	var doc: [String]
	var internalType: Int

	enum CodingKeys: String, CodingKey {
		case name = "name"
		case kind = "kind"
		case doc = "doc"
		case internalType = "internalType"
	}
}
```

PHP: <a name="php"></a>

```php
<?php


class EmbeddedType {
	public string $id;


	public function __construct(array $data) {
		$this->id = $data['id'];
		$this->data = $data['data'];
	}
}

class ExtendedScannedType extends ScannedType {
	public bool $anExtraField;


	public function __construct(array $data) {
		$this->anExtraField = $data['anExtraField'];
	}
}

class ScannedStruct {
	/**
	 * @var string[] $doc
	 */
	public array $doc;

	public string $name;

	/**
	 * @var bool[] $fields
	 */
	public array $fields;

	public int $internalType;


	public function __construct(array $data) {
		$this->doc = $data['doc'];
		$this->name = $data['name'];
		$this->fields = $data['fields'];
		$this->internalType = $data['internalType'];
	}
}

class ScannedType {
	public string $name;

	public string $kind;

	/**
	 * @var string[] $doc
	 */
	public array $doc;

	public int $internalType;


	public function __construct(array $data) {
		$this->name = $data['name'];
		$this->kind = $data['kind'];
		$this->doc = $data['doc'];
		$this->internalType = $data['internalType'];
	}
}
```

JSON: <a name="json"></a>

```json
{
  "enums": {},
  "types": {
    "EmbeddedType": {
      "doc": null,
      "name": "EmbeddedType",
      "fields": [
        {
          "name": "ID",
          "kind": "string",
          "tag": "`json:\"id\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        },
        {
          "name": "Data",
          "kind": "struct",
          "tag": "`json:\"data\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": [
            {
              "name": "Test",
              "kind": "string",
              "tag": "`json:\"test\"`",
              "doc": null,
              "importedEntity": null,
              "subFields": null,
              "internalType": 3
            }
          ],
          "internalType": 0
        }
      ],
      "internalType": 4
    },
    "ExtendedScannedType": {
      "doc": null,
      "name": "ExtendedScannedType",
      "fields": [
        {
          "name": "ScannedType",
          "kind": "struct",
          "tag": "",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 0
        },
        {
          "name": "AnExtraField",
          "kind": "bool",
          "tag": "`json:\"anExtraField\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        }
      ],
      "internalType": 4
    },
    "ScannedStruct": {
      "doc": null,
      "name": "ScannedStruct",
      "fields": [
        {
          "name": "Doc",
          "kind": "[]string",
          "tag": "`json:\"doc\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 5
        },
        {
          "name": "Name",
          "kind": "string",
          "tag": "`json:\"name\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        },
        {
          "name": "Fields",
          "kind": "[]bool",
          "tag": "`json:\"fields\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 5
        },
        {
          "name": "InternalType",
          "kind": "int",
          "tag": "`xml:\"internalType\" binding:\"required\" validation:\"required\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        }
      ],
      "internalType": 4
    },
    "ScannedType": {
      "doc": null,
      "name": "ScannedType",
      "fields": [
        {
          "name": "Name",
          "kind": "string",
          "tag": "`json:\"name\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        },
        {
          "name": "Kind",
          "kind": "string",
          "tag": "`json:\"kind\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        },
        {
          "name": "Doc",
          "kind": "[]string",
          "tag": "`json:\"doc\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 5
        },
        {
          "name": "InternalType",
          "kind": "int",
          "tag": "`json:\"internalType\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 3
        }
      ],
      "internalType": 4
    }
  }
}
```