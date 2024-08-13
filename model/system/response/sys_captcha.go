package response

type SysCaptchaResponse struct {
	CaptchaID     string `json:"captcha_id"`
	PicPath       string `json:"pic_path"`
	CaptchaLength int    `json:"captcha_length"`
	OpenCaptcha   bool   `json:"open_captcha"`
}
