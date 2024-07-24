package validation

var (
	MsgUnexpectedProperty                 = "unexpected property"
	MsgExpectedRFC3339DateTime            = "expected string to be RFC 3339 date-time"
	MsgExpectedRFC1123DateTime            = "expected string to be RFC 1123 date-time"
	MsgExpectedRFC3339Date                = "expected string to be RFC 3339 date"
	MsgExpectedRFC3339Time                = "expected string to be RFC 3339 time"
	MsgExpectedRFC5322Email               = "expected string to be RFC 5322 email: %v"
	MsgExpectedRFC5890Hostname            = "expected string to be RFC 5890 hostname"
	MsgExpectedRFC2673IPv4                = "expected string to be RFC 2673 ipv4"
	MsgExpectedRFC2373IPv6                = "expected string to be RFC 2373 ipv6"
	MsgExpectedRFC3986URI                 = "expected string to be RFC 3986 uri: %v"
	MsgExpectedRFC4122UUID                = "expected string to be RFC 4122 uuid: %v"
	MsgExpectedRFC6570URITemplate         = "expected string to be RFC 6570 uri-template"
	MsgExpectedRFC6901JSONPointer         = "expected string to be RFC 6901 json-pointer"
	MsgExpectedRFC6901RelativeJSONPointer = "expected string to be RFC 6901 relative-json-pointer"
	MsgExpectedRegexp                     = "expected string to be regex: %v"
	MsgExpectedMatchSchema                = "expected value to match at least one schema but matched none"
	MsgExpectedNotMatchSchema             = "expected value to not match schema"
	MsgExpectedBoolean                    = "expected boolean"
	MsgExpectedNumber                     = "expected number"
	MsgExpectedString                     = "expected string"
	MsgExpectedBase64String               = "expected string to be base64 encoded"
	MsgExpectedArray                      = "expected array"
	MsgExpectedObject                     = "expected object"
	MsgExpectedArrayItemsUnique           = "expected array items to be unique"
	MsgExpectedOneOf                      = "expected value to be one of \"%s\""
	MsgExpectedMinimumNumber              = "expected number >= %v"
	MsgExpectedExclusiveMinimumNumber     = "expected number > %v"
	MsgExpectedMaximumNumber              = "expected number <= %v"
	MsgExpectedExclusiveMaximumNumber     = "expected number < %v"
	MsgExpectedNumberBeMultipleOf         = "expected number to be a multiple of %v"
	MsgExpectedMinLength                  = "expected length >= %d"
	MsgExpectedMaxLength                  = "expected length <= %d"
	MsgExpectedBePattern                  = "expected string to be %s"
	MsgExpectedMatchPattern               = "expected string to match pattern %s"
	MsgExpectedMinItems                   = "expected array length >= %d"
	MsgExpectedMaxItems                   = "expected array length <= %d"
	MsgExpectedMinProperties              = "expected object with at least %d properties"
	MsgExpectedMaxProperties              = "expected object with at most %d properties"
	MsgExpectedRequiredProperty           = "expected required property %s to be present"
	MsgExpectedDependentRequiredProperty  = "expected property %s to be present when %s is present"
)
