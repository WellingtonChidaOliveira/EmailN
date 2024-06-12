package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrDataBase error = errors.New("error to save on database")