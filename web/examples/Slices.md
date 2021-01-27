# Slices

Below you can find some examples of how slices get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
var StringSlice = []string{
    "test",
    "test2",
}

var IntSlice = []int{
    123,
    234,
}

var FloatSlice = []float32{
    1.234,
    2.345,
}

var BoolSlice = []bool{
    true,
    false,
}

var AnySlice = []interface{}{
    true,
    "test2",
}
```

TypeScript: <a name="ts"></a>

```ts
export const AnySlice: any[] = [
    true,
    "test2"
];

export const BoolSlice: boolean[] = [
    true,
    false
];

export const FloatSlice: number[] = [
    1.234,
    2.345
];

export const IntSlice: number[] = [
    123,
    234
];

export const StringSlice: string[] = [
    "test",
    "test2"
];
```

Swift: <a name="swift"></a>

```swift
var AnySlice: [Any] = [
	true,
	"test2"
];

var BoolSlice: [Bool] = [
	true,
	false
];

var FloatSlice: [Float] = [
	1.234,
	2.345
];

var IntSlice: [Int] = [
	123,
	234
];

var StringSlice: [String] = [
	"test",
	"test2"
];
```

PHP: <a name="php"></a>

```php
<?php

/**
 * @const array AnySlice
 */
const AnySlice = [
	true,
	"test2"
];

/**
 * @const bool[] BoolSlice
 */
const BoolSlice = [
	true,
	false
];

/**
 * @const float[] FloatSlice
 */
const FloatSlice = [
	1.234,
	2.345
];

/**
 * @const int[] IntSlice
 */
const IntSlice = [
	123,
	234
];

/**
 * @const string[] StringSlice
 */
const StringSlice = [
	"test",
	"test2"
];
```

JSON: <a name="json"></a>

```json
{
  "enums": {
    "AnySlice": {
      "name": "AnySlice",
      "kind": "[]interface{}",
      "value": [
        "true",
        "\"test2\""
      ],
      "doc": null,
      "internalType": 5
    },
    "BoolSlice": {
      "name": "BoolSlice",
      "kind": "[]bool",
      "value": [
        "true",
        "false"
      ],
      "doc": null,
      "internalType": 5
    },
    "FloatSlice": {
      "name": "FloatSlice",
      "kind": "[]float32",
      "value": [
        "1.234",
        "2.345"
      ],
      "doc": null,
      "internalType": 5
    },
    "IntSlice": {
      "name": "IntSlice",
      "kind": "[]int",
      "value": [
        "123",
        "234"
      ],
      "doc": null,
      "internalType": 5
    },
    "StringSlice": {
      "name": "StringSlice",
      "kind": "[]string",
      "value": [
        "\"test\"",
        "\"test2\""
      ],
      "doc": null,
      "internalType": 5
    }
  },
  "types": {}
}
```