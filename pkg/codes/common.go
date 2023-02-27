package codes

import "github.com/richkeyu/gocommons/perrors"

const (
	CodeInfoSuccess  = 1
	CodeErrorUnknown = -1
	CodeNoNeedUpdate = 20001

	CodeErrorData                   = 10000
	CodeErrorTimeout                = 10001
	CodeErrorForm                   = 10002
	CodeErrorParameter              = 10003
	CodeErrorValidation             = 10004
	CodeErrorPermissionDenied       = 10006
	CodeErrorAccountNoExist         = 10007
)

var InfoSuccess = perrors.NewError(CodeInfoSuccess, "success")          //  成功
var ErrorUnknown = perrors.NewError(CodeErrorUnknown, "unknown error")  //  未知错误
var NoNeedUpdate = perrors.NewError(CodeNoNeedUpdate, "no need update") //  无需更新

var ErrorTimeout = perrors.NewError(CodeErrorTimeout, "timeout error")                                                          //  超时
var ErrorForm = perrors.NewError(CodeErrorForm, "form error")                                                                   //  表单错误
var ErrorData = perrors.NewError(CodeErrorData, "data error")                                                                   //  数据错误
var ErrorParameter = perrors.NewError(CodeErrorParameter, "parameter error")                                                    //  参数错误
var ErrorValidation = perrors.NewError(CodeErrorValidation, "validation error")                                                 //  校验错误
var ErrorPermissionDenied = perrors.NewError(CodeErrorPermissionDenied, "permission denied")                                    //  权限不足
var ErrorAccountNoExist = perrors.NewError(CodeErrorAccountNoExist, "account not exist")                                        //  账号不存在
