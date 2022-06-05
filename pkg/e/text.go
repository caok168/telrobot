package e

var errorText = map[int] string{
	ErrorOK:                    "success",
	ErrorUnknown:               "Unknown error",
	ErrorAuthTokenRequired:     "Token is required",
	ErrorAuthCheckTokenFail:    "Check token fail",
	ErrorAuthCheckTokenTimeout: "Check token timeout",
	ErrorTooManyRequest:        "Too many request",
}
