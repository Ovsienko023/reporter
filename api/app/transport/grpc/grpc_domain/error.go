package grpc_domain

//import (
//	"google.golang.org/genproto/googleapis/rpc/errdetails"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//)
//
//func HandleAppError(err error) error {
//	if err == nil {
//		return nil
//	}
//
//	var appError *errcore.AppError
//
//	if !errors.As(err, &appError) {
//		return status.New(codes.Internal, err.Error()).Err()
//	}
//
//	var code codes.Code
//
//	switch appError.ErrorType {
//	case errcore.AccessType:
//		code = codes.PermissionDenied
//	case errcore.ObjectNotFoundType:
//		code = codes.NotFound
//	case errcore.ObjectAlreadyExistType:
//		code = codes.AlreadyExists
//	case errcore.ObjectRequiredType, errcore.ObjectMissingType, errcore.InvalidRequestType:
//		code = codes.InvalidArgument
//	default:
//		code = codes.Internal
//	}
//
//	st := status.New(code, appError.Message)
//
//	if appError.Details != nil {
//		details := HandleAppErrorDetails(appError)
//
//		statusWithErr, errs := st.WithDetails(details)
//		if errs == nil {
//			return statusWithErr.Err()
//		}
//	}
//
//	return st.Err()
//}
//
//func HandleAppErrorDetails(appError *errcore.AppError) *errdetails.BadRequest {
//	if appError == nil {
//		return nil
//	}
//
//	details := &errdetails.BadRequest{}
//
//	for _, errorDetail := range appError.Details {
//		switch t := errorDetail.(type) {
//		case *errcore.ValidationDetail:
//			details = &errdetails.BadRequest{
//				FieldViolations: make([]*errdetails.BadRequest_FieldViolation, 0, len(t.Errors)),
//			}
//
//			for _, field := range t.Errors {
//				details.FieldViolations = append(details.FieldViolations, &errdetails.BadRequest_FieldViolation{
//					Field:       field.Field,
//					Description: field.Description,
//				})
//			}
//		}
//	}
//
//	return details
//}
