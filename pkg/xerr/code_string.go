// Code generated by "stringer -type Code -linecomment"; DO NOT EDIT.

package xerr

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrSuccess-0]
	_ = x[ErrUnknown-100001]
	_ = x[ErrParam-100002]
	_ = x[ErrHttpClient-100003]
	_ = x[ErrLoginExpired-100101]
	_ = x[ErrAccountLoggedOut-100102]
	_ = x[ErrTokenInvalid-110001]
	_ = x[ErrCaptchaInvalid-110002]
	_ = x[ErrCaptchaWrong-110003]
	_ = x[ErrDeviceIDInconsistent-110004]
	_ = x[ErrPhoneNumWrong-110005]
	_ = x[ErrSendLimit-110006]
	_ = x[ErrCodeWrong-110007]
	_ = x[ErrCodeWrongThreeTimes-110008]
	_ = x[ErrUserNotFound-110101]
	_ = x[ErrElectricityTokenInvalid-110102]
	_ = x[ErrElectricityBindNotFound-110103]
}

const (
	_Code_name_0 = "Success"
	_Code_name_1 = "服务异常参数错误HTTP客户端请求错误"
	_Code_name_2 = "登录已过期账号被登出"
	_Code_name_3 = "Token无效图片验证码已失效图片验证码错误deviceId不一致手机号格式错误短信发送超限手机验证码错误, 错误3次将锁定15分钟手机验证码错误3次, 账号锁定15分钟"
	_Code_name_4 = "用户不存在电费Token无效未找到电费绑定信息"
)

var (
	_Code_index_1 = [...]uint8{0, 12, 24, 49}
	_Code_index_2 = [...]uint8{0, 15, 30}
	_Code_index_3 = [...]uint8{0, 11, 35, 56, 73, 94, 112, 162, 209}
	_Code_index_4 = [...]uint8{0, 15, 32, 59}
)

func (i Code) String() string {
	switch {
	case i == 0:
		return _Code_name_0
	case 100001 <= i && i <= 100003:
		i -= 100001
		return _Code_name_1[_Code_index_1[i]:_Code_index_1[i+1]]
	case 100101 <= i && i <= 100102:
		i -= 100101
		return _Code_name_2[_Code_index_2[i]:_Code_index_2[i+1]]
	case 110001 <= i && i <= 110008:
		i -= 110001
		return _Code_name_3[_Code_index_3[i]:_Code_index_3[i+1]]
	case 110101 <= i && i <= 110103:
		i -= 110101
		return _Code_name_4[_Code_index_4[i]:_Code_index_4[i+1]]
	default:
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
