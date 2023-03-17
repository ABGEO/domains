package forms

type ProtectedForm struct {
	RecaptchaToken string `form:"recaptchaToken" binding:"required"`
}
