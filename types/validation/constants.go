package validation

// key is struct field, value is error message
type FormErrors map[string]string

const REQUIRE_FIELD_MESSAGE string = "This field is required"
const VALID_EMAIL_MESSAGE string = "A valid email address is required"
const PASSWORD_MUST_MATCH_MESSAGE string = "Password and Confirm Password must match"

const MINIMUM_PASSWORD_LENGTH int = 8
const PASSWORD_MINIMUM_LENGTH_MESSAGE string = "Your password must be at least 8 characters"
