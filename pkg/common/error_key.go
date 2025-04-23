package common

type ResponseCode string

const (
	ErrUserIDRequired   = "user_id_required"
	ErrUserNotFound     = "user_not_found"
	ErrInvalidTimestamp = "invalid_timestamp"
	ErrHMACMismatch     = "hmac_mismatch"
	ErrInternal         = "internal_error"
)
