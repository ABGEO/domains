package controllers

import (
	"net/http"

	"github.com/abgeo/domains/config"
	"github.com/abgeo/domains/forms"
	"github.com/abgeo/domains/services"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (controller DefaultController) Index(ctx *gin.Context) {
	conf := config.GetConfig()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"domain":           ctx.Request.Host,
		"recaptchaSiteKey": conf.Recaptcha.SiteKey,
		"gaCode":           conf.GACode,
		"contactEmail":     conf.Email.From,
	})
}

func (controller DefaultController) Submit(ctx *gin.Context) {
	conf := config.GetConfig()
	mailerService := services.NewMailerService(conf)
	recaptchaService := services.NewRecaptchaService(conf)

	var formData forms.Offer
	if err := ctx.Bind(&formData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	if err := recaptchaService.Validate(ctx.Request.Host, formData.RecaptchaToken); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	if err := mailerService.Send(formData.Email, "offer", formData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
