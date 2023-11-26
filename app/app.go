package app

import (
	"context"

	style "github.com/a-yee/spot/ui/Style"
	"github.com/a-yee/spot/ui/keymap"
	api "github.com/zmb3/spotify/v2"
)

type AppInfo struct {
	ctx    context.Context
	API    *api.Client
	Width  int
	Height int
	Style  *style.Style
	KeyMap *keymap.KeyMap
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
		Style:  style.DefaultStyles(),
		KeyMap: keymap.DefaultKeyMap(),
	}
}

// Set the width and height of the component
func (a *AppInfo) SetSize(width, height int) {
	a.Width = width
	a.Height = height
}
