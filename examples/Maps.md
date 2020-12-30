# Maps

Below you can find some examples of how maps get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
var StringStringMap = map[string]string{ //nolint:gochecknoglobals
    "test":  "123",
    "test2": "234",
}

var IntIntMap = map[int]int{ //nolint:gochecknoglobals
    0: 123,
    1: 234,
}

var IntStringMap = map[int]string{ //nolint:gochecknoglobals
    0: "test",
    1: "test1",
}

var StringIntMap = map[string]int{ //nolint:gochecknoglobals
    "test":  123,
    "test1": 234,
}

var StringAnyMap = map[string]interface{}{ //nolint:gochecknoglobals
    "test a string": "test",
    "test a int":    123,
    "test a bool":   false,
}

var StringBoolMap = map[string]bool{ //nolint:gochecknoglobals
    "godmt": true,
    "typex": false,
}
```

TypeScript: <a name="ts"></a>

```ts
export const IntIntMap: Record<number, number> = {
    0: 123,
    1: 234
};

export const IntStringMap: Record<number, string> = {
    0: "test",
    1: "test1"
};

export const StringAnyMap: Record<string, any> = {
    "test a string": "test",
    "test a int": 123,
    "test a bool": false
};

export const StringBoolMap: Record<string, boolean> = {
    "godmt": true,
    "typex": false
};

export const StringIntMap: Record<string, number> = {
    "test": 123,
    "test1": 234
};

export const StringStringMap: Record<string, string> = {
    "test": "123",
    "test2": "234"
};

```

Swift: <a name="swift"></a>

```swift
let IntIntMap: Dictionary<Int, Int> = [
	0: 123,
	1: 234
]

let IntStringMap: Dictionary<Int, String> = [
	0: "test",
	1: "test1"
]

let StringAnyMap: Dictionary<String, Any> = [
	"test a string": "test",
	"test a int": 123,
	"test a bool": false
]

let StringBoolMap: Dictionary<String, Bool> = [
	"godmt": true,
	"typex": false
]

let StringIntMap: Dictionary<String, Int> = [
	"test": 123,
	"test1": 234
]

let StringStringMap: Dictionary<String, String> = [
	"test": "123",
	"test2": "234"
]
```

PHP: <a name="php"></a>

```php
<?php

/**
 * @const array IntIntMap
 */
const IntIntMap = [
	0 => 123,
	1 => 234
];

/**
 * @const array IntStringMap
 */
const IntStringMap = [
	0 => "test",
	1 => "test1"
];

/**
 * @const array StringAnyMap
 */
const StringAnyMap = [
	"test a string" => "test",
	"test a int" => 123,
	"test a bool" => false
];

/**
 * @const array StringBoolMap
 */
const StringBoolMap = [
	"typex" => false,
	"godmt" => true
];

/**
 * @const array StringIntMap
 */
const StringIntMap = [
	"test" => 123,
	"test1" => 234
];

/**
 * @const array StringStringMap
 */
const StringStringMap = [
	"test" => "123",
	"test2" => "234"
];
```

JSON: <a name="json"></a>

```json
{
  "enums": {
    "IntIntMap": {
      "name": "IntIntMap",
      "kind": "map[int]int",
      "value": {
        "0": "123",
        "1": "234"
      },
      "doc": null,
      "internalType": 1
    },
    "IntStringMap": {
      "name": "IntStringMap",
      "kind": "map[int]string",
      "value": {
        "0": "\"test\"",
        "1": "\"test1\""
      },
      "doc": null,
      "internalType": 1
    },
    "StringAnyMap": {
      "name": "StringAnyMap",
      "kind": "map[string]interface{}",
      "value": {
        "\"test a bool\"": "false",
        "\"test a int\"": "123",
        "\"test a string\"": "\"test\""
      },
      "doc": null,
      "internalType": 1
    },
    "StringBoolMap": {
      "name": "StringBoolMap",
      "kind": "map[string]bool",
      "value": {
        "\"godmt\"": "true",
        "\"typex\"": "false"
      },
      "doc": null,
      "internalType": 1
    },
    "StringIntMap": {
      "name": "StringIntMap",
      "kind": "map[string]int",
      "value": {
        "\"test\"": "123",
        "\"test1\"": "234"
      },
      "doc": null,
      "internalType": 1
    },
    "StringStringMap": {
      "name": "StringStringMap",
      "kind": "map[string]string",
      "value": {
        "\"test\"": "\"123\"",
        "\"test2\"": "\"234\""
      },
      "doc": null,
      "internalType": 1
    }
  },
  "types": {}
}
```