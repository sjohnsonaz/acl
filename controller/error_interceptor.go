package controller

import (
	"context"
	"errors"

	"github.com/cardboardrobots/baseerror"
	"google.golang.org/grpc/codes"
)

func GetStatusCode(err error) codes.Code {
	if errors.Is(err, context.Canceled) {
		return codes.Canceled
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return codes.DeadlineExceeded
	}
	if errors.Is(err, baseerror.ErrAlreadyExists) {
		return codes.AlreadyExists
	}
	if errors.Is(err, baseerror.ErrFailedPrecondition) {
		return codes.FailedPrecondition
	}
	if errors.Is(err, baseerror.ErrInvalidArgument) {
		return codes.InvalidArgument
	}
	if errors.Is(err, baseerror.ErrNotFound) {
		return codes.NotFound
	}
	if errors.Is(err, baseerror.ErrOutOfRange) {
		return codes.OutOfRange
	}
	if errors.Is(err, baseerror.ErrPermissionDenied) {
		return codes.PermissionDenied
	}
	if errors.Is(err, baseerror.ErrUnauthenticated) {
		return codes.Unauthenticated
	}
	return codes.Unknown
}
