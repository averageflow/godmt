# Complex Structures

Below you can find some examples of how complicated structures get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
// ComplicatedMaps is a class that deserves a doc comment
type ComplicatedMaps struct {
    // StrangeMapOne is a complicated type
    StrangeMapOne map[string]map[string]int `json:"strangeMapOne"`
    // StrangeSliceOfMaps is also complicated type
    StrangeSliceOfMaps []map[int]string         `json:"strangeSliceOfMaps"`
    SimpleMap          map[string]int           `json:"simpleMap"`
    MapOfSlices        map[string][]int         `json:"tricky"`
    Insane             map[int][]map[string]int `json:"insane"`
}
```

TypeScript: <a name="ts"></a>

```ts
export interface ComplicatedMaps {
    // StrangeMapOne is a complicated type
    strangeMapOne: Record<string, Record<string, number>>;
    // StrangeSliceOfMaps is also complicated type
    strangeSliceOfMaps: Record<number, string>[];
    simpleMap: Record<string, number>;
    tricky: Record<string, number[]>;
    insane: Record<number, Record<string, number>[]>;
}
```

Swift: <a name="swift"></a>

```swift
struct ComplicatedMaps: Decodable {
	// StrangeMapOne is a complicated type
	var strangeMapOne: Dictionary<String, Dictionary<String, Int>>
	// StrangeSliceOfMaps is also complicated type
	var strangeSliceOfMaps: [Dictionary<Int, String>]
	var simpleMap: Dictionary<String, Int>
	var tricky: Dictionary<String, [Int]>
	var insane: Dictionary<Int, [Dictionary<String, Int>]>

	enum CodingKeys: String, CodingKey {
		case strangeMapOne = "strangeMapOne"
		case strangeSliceOfMaps = "strangeSliceOfMaps"
		case simpleMap = "simpleMap"
		case tricky = "tricky"
		case insane = "insane"
	}
}
```

PHP: <a name="php"></a>

```php
<?php


class ComplicatedMaps {
	/**
	 * StrangeMapOne is a complicated type
	 * @var array $strangeMapOne
	 */
	public array $strangeMapOne;

	/**
	 * StrangeSliceOfMaps is also complicated type
	 * @var array $strangeSliceOfMaps
	 */
	public array $strangeSliceOfMaps;

	/**
	 * @var array $simpleMap
	 */
	public array $simpleMap;

	/**
	 * @var array $tricky
	 */
	public array $tricky;

	/**
	 * @var array $insane
	 */
	public array $insane;


	public function __construct(array $data) {
		$this->strangeMapOne = $data['strangeMapOne'];
		$this->strangeSliceOfMaps = $data['strangeSliceOfMaps'];
		$this->simpleMap = $data['simpleMap'];
		$this->tricky = $data['tricky'];
		$this->insane = $data['insane'];
	}
}
```

JSON: <a name="json"></a>

```json
{
  "enums": {},
  "types": {
    "ComplicatedMaps": {
      "doc": null,
      "name": "ComplicatedMaps",
      "fields": [
        {
          "name": "StrangeMapOne",
          "kind": "map[string]map[string]int",
          "tag": "`json:\"strangeMapOne\"`",
          "doc": [
            "// StrangeMapOne is a complicated type"
          ],
          "importedEntity": null,
          "subFields": null,
          "internalType": 1
        },
        {
          "name": "StrangeSliceOfMaps",
          "kind": "[]map[int]string",
          "tag": "`json:\"strangeSliceOfMaps\"`",
          "doc": [
            "// StrangeSliceOfMaps is also complicated type"
          ],
          "importedEntity": null,
          "subFields": null,
          "internalType": 5
        },
        {
          "name": "SimpleMap",
          "kind": "map[string]int",
          "tag": "`json:\"simpleMap\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1
        },
        {
          "name": "MapOfSlices",
          "kind": "map[string][]int",
          "tag": "`json:\"tricky\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1
        },
        {
          "name": "Insane",
          "kind": "map[int][]map[string]int",
          "tag": "`json:\"insane\"`",
          "doc": null,
          "importedEntity": null,
          "subFields": null,
          "internalType": 1
        }
      ],
      "internalType": 4
    }
  }
}
```