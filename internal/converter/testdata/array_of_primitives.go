package testdata

const ArrayOfPrimitives = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "luckyNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyBigNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "string"
                    },
                    {
                        "type": "null"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "keyWords": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "string"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "big_number": {
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`

const ArrayOfPrimitivesDouble = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "description": {
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "string"
                }
            ]
        },
        "luckyNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "integer"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "luckyBigNumbers": {
            "items": {
                "oneOf": [
                    {
                        "type": "string"
                    },
                    {
                        "type": "null"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "keyWords": {
            "items": {
                "oneOf": [
                    {
                        "type": "null"
                    },
                    {
                        "type": "string"
                    }
                ]
            },
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "array"
                }
            ]
        },
        "big_number": {
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        },
        "bigNumber": {
            "oneOf": [
                {
                    "type": "string"
                },
                {
                    "type": "null"
                }
            ]
        }
    },
    "additionalProperties": true,
    "oneOf": [
        {
            "type": "null"
        },
        {
            "type": "object"
        }
    ]
}`
