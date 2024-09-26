package response

type SysCaptchaResponse struct {
	CaptchaID     string `json:"captchaID"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
	OpenCaptcha   bool   `json:"openCaptcha"`
}
