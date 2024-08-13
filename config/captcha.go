package config

type Captcha struct {
	KeyLong            int `json:"key_long" yaml:"key_long" mapstructure:"key_long"`                                        // 验证码长度
	ImgWidth           int `json:"img_width" yaml:"img_width" mapstructure:"img_width"`                                     // 验证码宽度
	ImgHeight          int `json:"img_height" yaml:"img_height" mapstructure:"img_height"`                                  // 验证码高度
	OpenCaptcha        int `json:"open_captcha" yaml:"open_captcha" mapstructure:"open_captcha"`                            //防爆验证码开启次数 0代表每次登录都需要验证码 其他数字代表错误密码次数 例3代表错误3次之后出现验证码
	OpenCaptchaTimeOut int `json:"open_captcha_time_out" yaml:"open_captcha_time_out" mapstructure:"open_captcha_time_out"` // 防爆验证码超时时间 单位 s(秒)
}
