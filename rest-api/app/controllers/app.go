// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"aahframework.org/aah.v0"
	"aahframework.org/examples/rest-api/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application root API endpoint.
func (c *AppController) Index() {
	c.Reply().Ok().JSON(models.Greet{
		Message: "Welcome to aah framework - REST API Services example",
	})
}
