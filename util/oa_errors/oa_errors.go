package oa_errors

import (
	"Oasyce-backend/api/kitex/kitex_gen/common"
)

// Error defines interface that provide Error elements
type Error interface {
	GetHTTPCode() int32
	GetCode() string
	GetMessage() string
	GetData() map[string]string
}

// Deprecated:ToCommonError ...
func ToCommonError(errGen Error) *common.Error {
	return &common.Error{
		HTTPCode: errGen.GetHTTPCode(),
		Code:     errGen.GetCode(),
		Message:  errGen.GetMessage(),
		Data:     errGen.GetData(),
	}
}

type errorBase struct {
	httpCode int32
	code     string
	message  string
	data     map[string]string

	dataPreset map[string]string
}

// GetHTTPCode returns http code of the error
func (e *errorBase) GetHTTPCode() int32 {
	return e.httpCode
}

// GetCode returns code of the error
func (e *errorBase) GetCode() string {
	return e.code
}

// GetMessage returns message of the error
func (e *errorBase) GetMessage() string {
	return e.message
}

// GetData returns data map of the error
func (e *errorBase) GetData() map[string]string {
	return func(mObj ...map[string]string) map[string]string {
		var ret map[string]string
		for _, m := range mObj {
			for k, v := range m {
				if ret == nil {
					ret = make(map[string]string)
				}
				ret[k] = v
			}
		}
		return ret
	}(e.dataPreset, e.data)
}

// ResponseWithCommonError is the common interface for the IDL result type that throws commondto.Error.
// It can be converted from response of type interface{} and set the `Err` field.
// Usually, it should be used in middleware to return commondto.Error.
// Usage: https://bytedance.sg.feishu.cn/wiki/wikcnqudvX62meLnzsz86Wn6ALg
//
// e.g.
// // vecode/api/kitex/kitex_gen/vecode/vecode.go
//
//		type RepositoryServiceCreateRepositoryResult struct {
//		  Success *repository.RepositoryResponse `thrift:"success,0" json:"success,omitempty"`
//	      Err     *commondto.Error               `thrift:"err,1" json:"err,omitempty"`
//		}
type ResponseWithCommonError interface {
	SetErr(err *common.Error)
	GetErr() (v *common.Error)
}
