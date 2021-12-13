package test 

import {
	"context"
	_ "embed"
	"encoding/xml"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/Nsadow311/stranger/common"
	"github.com/Nsadow311/stranger/common/cplogs"
	"github.com/Nsadow311/stranger/web"
	"github.com/jonas747/discordgo/v2"
	
}

//go:embed assets/test.html
var PageHTML string

func (p *Plugin) InitWeb() {
	web.AddHTMLTemplate("test/assets/test.html", PageHTML)
	web.AddSidebarItem(web.SidebarCategoryFeeds, &web.SidebarItem{
		Name: "Youtube",
		URL:  "youtube",
		Icon: "fab fa-youtube",
	})