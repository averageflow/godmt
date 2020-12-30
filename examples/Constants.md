# Constants

Below you can find some examples of how constants get translated to other languages.

- [Go](#go)
- [TypeScript](#ts)
- [Swift](#swift)
- [PHP](#php)
- [JSON](#json)

Go: <a name="go"></a>

```go
const (
	ExampleInt    = 1
	ExampleFloat  = 3.141516
	ExampleString = "test"
	ExampleBool   = true
)

const (
	// ExampleIntWithComment represents an integer with a comment line above it.
	ExampleIntWithComment = 2
	// ExampleFloatWithComment represents a float with a comment line above it.
	ExampleFloatWithComment = 4.141516
	// ExampleStringWithComment represents a string with a comment line above it.
	ExampleStringWithComment = "test2"
	// ExampleBoolWithComment represents a bool with a comment line above it.
	ExampleBoolWithComment = false
)
```

TypeScript: <a name="ts"></a>

```ts
export const ExampleBool: boolean = true;

// ExampleBoolWithComment represents a bool with a comment line above it.
export const ExampleBoolWithComment: boolean = false;

export const ExampleFloat: number = 3.141516;

// ExampleFloatWithComment represents a float with a comment line above it.
export const ExampleFloatWithComment: number = 4.141516;

export const ExampleInt: number = 1;

// ExampleIntWithComment represents an integer with a comment line above it.
export const ExampleIntWithComment: number = 2;

export const ExampleString: string = "test";

// ExampleStringWithComment represents a string with a comment line above it.
export const ExampleStringWithComment: string = "test2";
```

Swift: <a name="swift"></a>

```swift
let ExampleBool: Bool = true

// ExampleBoolWithComment represents a bool with a comment line above it.
let ExampleBoolWithComment: Bool = false

let ExampleFloat: Float = 3.141516

// ExampleFloatWithComment represents a float with a comment line above it.
let ExampleFloatWithComment: Float = 4.141516

let ExampleInt: Int = 1

// ExampleIntWithComment represents an integer with a comment line above it.
let ExampleIntWithComment: Int = 2

let ExampleString: String = "test"

// ExampleStringWithComment represents a string with a comment line above it.
let ExampleStringWithComment: String = "test2"
```

PHP: <a name="php"></a>

```php
<?php

/**
 * @const ExampleBool bool
 */
const ExampleBool = true;

/**
 * ExampleBoolWithComment represents a bool with a comment line above it.
 * @const ExampleBoolWithComment bool
 */
const ExampleBoolWithComment = false;

/**
 * @const ExampleFloat float
 */
const ExampleFloat = 3.141516;

/**
 * ExampleFloatWithComment represents a float with a comment line above it.
 * @const ExampleFloatWithComment float
 */
const ExampleFloatWithComment = 4.141516;

/**
 * @const ExampleInt int
 */
const ExampleInt = 1;

/**
 * ExampleIntWithComment represents an integer with a comment line above it.
 * @const ExampleIntWithComment int
 */
const ExampleIntWithComment = 2;

/**
 * @const ExampleString string
 */
const ExampleString = "test";

/**
 * ExampleStringWithComment represents a string with a comment line above it.
 * @const ExampleStringWithComment string
 */
const ExampleStringWithComment = "test2";
```

JSON: <a name="json"></a>

```json
{
	"enums": {
		"ExampleBool": {
			"name": "ExampleBool",
			"kind": "bool",
			"value": "true",
			"doc": null,
			"internalType": 3
		},
		"ExampleBoolWithComment": {
			"name": "ExampleBoolWithComment",
			"kind": "bool",
			"value": "false",
			"doc": [
				"// ExampleBoolWithComment represents a bool with a comment line above it."
			],
			"internalType": 3
		},
		"ExampleFloat": {
			"name": "ExampleFloat",
			"kind": "float64",
			"value": "3.141516",
			"doc": null,
			"internalType": 3
		},
		"ExampleFloatWithComment": {
			"name": "ExampleFloatWithComment",
			"kind": "float64",
			"value": "4.141516",
			"doc": [
				"// ExampleFloatWithComment represents a float with a comment line above it."
			],
			"internalType": 3
		},
		"ExampleInt": {
			"name": "ExampleInt",
			"kind": "int64",
			"value": "1",
			"doc": null,
			"internalType": 3
		},
		"ExampleIntWithComment": {
			"name": "ExampleIntWithComment",
			"kind": "int64",
			"value": "2",
			"doc": [
				"// ExampleIntWithComment represents an integer with a comment line above it."
			],
			"internalType": 3
		},
		"ExampleString": {
			"name": "ExampleString",
			"kind": "string",
			"value": "\"test\"",
			"doc": null,
			"internalType": 3
		},
		"ExampleStringWithComment": {
			"name": "ExampleStringWithComment",
			"kind": "string",
			"value": "\"test2\"",
			"doc": [
				"// ExampleStringWithComment represents a string with a comment line above it."
			],
			"internalType": 3
		}
	},
	"types": {}
}
```