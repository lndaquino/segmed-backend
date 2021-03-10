package errors

import "errors"

// ErrFileNotFound is used when no file record was found in database
var ErrFileNotFound = errors.New("file not found")
