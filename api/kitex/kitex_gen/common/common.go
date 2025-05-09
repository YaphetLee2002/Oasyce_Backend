// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package common

import (
	"fmt"
)

type Error struct {
	HTTPCode int32             `thrift:"HTTPCode,1" frugal:"1,default,i32" json:"HTTPCode"`
	Code     string            `thrift:"Code,2" frugal:"2,default,string" json:"Code"`
	Message  string            `thrift:"Message,3" frugal:"3,default,string" json:"Message"`
	Data     map[string]string `thrift:"Data,4" frugal:"4,default,map<string:string>" json:"Data,omitempty"`
}

func NewError() *Error {
	return &Error{}
}

func (p *Error) InitDefault() {
}

func (p *Error) GetHTTPCode() (v int32) {
	return p.HTTPCode
}

func (p *Error) GetCode() (v string) {
	return p.Code
}

func (p *Error) GetMessage() (v string) {
	return p.Message
}

func (p *Error) GetData() (v map[string]string) {
	return p.Data
}
func (p *Error) SetHTTPCode(val int32) {
	p.HTTPCode = val
}
func (p *Error) SetCode(val string) {
	p.Code = val
}
func (p *Error) SetMessage(val string) {
	p.Message = val
}
func (p *Error) SetData(val map[string]string) {
	p.Data = val
}

func (p *Error) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Error(%+v)", *p)
}
func (p *Error) Error() string {
	return p.String()
}

var fieldIDToName_Error = map[int16]string{
	1: "HTTPCode",
	2: "Code",
	3: "Message",
	4: "Data",
}
