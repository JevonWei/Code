package controllers

import (
	"github.com/JevonWei/gocmdb/server/controllers/auth"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare() {
	c.Layout = "layouts/base.html"
}
