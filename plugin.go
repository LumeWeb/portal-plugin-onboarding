package plugin

import (
	"go.lumeweb.com/portal-plugin-onboarding/build"
	"go.lumeweb.com/portal-plugin-onboarding/internal"
	"go.lumeweb.com/portal/core"
	portal_plugin_onboarding "go.lumeweb.com/web/go/portal-plugin-onboarding"
)

func GetPluginInfo() core.PluginInfo {
	return core.PluginInfo{
		ID:         internal.PLUGIN_NAME,
		Version:    build.GetInfo(),
		WebBundles: core.NewWebBundles(core.NewWebBundle(portal_plugin_onboarding.GetFS())),
	}
}

func init() {
	core.RegisterPlugin(GetPluginInfo())
}
