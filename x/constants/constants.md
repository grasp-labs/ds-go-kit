# Constants

Shared constants, validation patterns, and compiled regular expressions.

## Validation Patterns

| Constant                        | Description                                                                                                                                                                                                                                 |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| RegPatternNotAlphaNumUnderscore | Match any character that is not a lowercase letter, digit, or underscore.                                                                                                                                                                   |
| RegPatternStringSearch          | Case-insensitive match for letters, digits, period, comma, space, underscore, and hyphen.                                                                                                                                                   |
| RegPatternAWSS3BucketName       | Validate an AWS S3 bucket name (lowercase letters, digits, hyphens).                                                                                                                                                                        |
| RegPatternAWSRegion             | Validate an AWS region format (e.g. eu-west-1).                                                                                                                                                                                             |
| RegPatternAWSS3Prefix           | Validate an AWS S3 prefix (zero or more path segments ending with /).                                                                                                                                                                       |
| RegPatternDPName                | Name attributes can only contain a-Z, norwegian special letters, 0-9 and whitespace                                                                                                                                                         |
| RegPatternDPServiceName         | Service name can only contain a-Z characters.                                                                                                                                                                                               |
| RegPatternProductID             | Segments 1–3: letters/digits/hyphens, must start/end alphanumeric (no leading/trailing '-') Version: optional v/V + 3 or 4 integer parts (e.g., 1.0.0 or v1.0.0.0)                                                                          |
| RegPatternTopicID               | Kafka topic format: {domain}.{subdomain}.{service}.{[optional]resource}.{action}.v{version} - Tokens: lowercase letters, digits, hyphens; must start alphanumeric. MAJOR >= 1 (disallows v0). Tokens: [a-z0-9][a-z0-9-]\* - prefixed with v |

Each pattern has a corresponding compiled `RegExp*` var ready for use.

## Default Values

| Constant        | Value     |
| --------------- | --------- |
| DefaultLanguage | Norwegian |
