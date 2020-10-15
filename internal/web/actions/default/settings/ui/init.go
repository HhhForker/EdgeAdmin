package ui

import (
	"github.com/TeaOSLab/EdgeAdmin/internal/web/actions/default/settings/settingutils"
	"github.com/TeaOSLab/EdgeAdmin/internal/web/helpers"
	"github.com/iwind/TeaGo"
)

func init() {
	TeaGo.BeforeStart(func(server *TeaGo.Server) {
		server.
			Helper(helpers.NewUserMustAuth()).
			Helper(settingutils.NewHelper("ui")).
			Prefix("/settings/ui").
			Get("", new(IndexAction)).
			GetPost("/updateHTTPPopup", new(UpdateHTTPPopupAction)).
			GetPost("/updateHTTPSPopup", new(UpdateHTTPSPopupAction)).
			EndAll()
	})
}
