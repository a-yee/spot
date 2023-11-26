package app

import (
	"context"

	api "github.com/zmb3/spotify/v2"
)

type AppInfo struct {
	ctx    context.Context
	API    *api.Client
	Width  int
	Height int
}

func NewAppInfo(
	ctx context.Context,
	apiClient *api.Client,
	width,
	height int) AppInfo {

	if ctx == nil {
		ctx = context.TODO()
	}

	return AppInfo{
		ctx:    ctx,
		API:    apiClient,
		Width:  width,
		Height: height,
	}
}
