package api

import "github.com/twelvet-s/gins/api/system"

type BuildApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var BuildGroup = new(BuildApiGroup)
