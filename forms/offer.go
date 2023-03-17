package forms

type Offer struct {
	ProtectedForm
	Name    string `form:"name" binding:"required"`
	Surname string `form:"surname" binding:"required"`
	Email   string `form:"email" binding:"required"`
	Phone   string `form:"phone"`
	Message string `form:"message" binding:"required"`
}
