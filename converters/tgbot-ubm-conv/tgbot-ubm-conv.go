package main

import (
	"github.com/BurntSushi/toml"
	"github.com/projectriri/bot-gateway/plugin"
	"github.com/projectriri/bot-gateway/router"
	"strings"
)

var (
	BuildTag      string
	BuildDate     string
	GitCommitSHA1 string
	GitTag        string
)

type Plugin struct{}

var manifest = plugin.Manifest{
	BasicInfo: plugin.BasicInfo{
		Name:    "tgbot-ubm-conv",
		Author:  "Project Riri Staff",
		Version: "v0.1",
		License: "MIT",
		URL:     "https://github.com/projectriri/bot-gateway/converters/tgbot-ubm-conv",
	},
	BuildInfo: plugin.BuildInfo{
		BuildTag:      BuildTag,
		BuildDate:     BuildDate,
		GitCommitSHA1: GitCommitSHA1,
		GitTag:        GitTag,
	},
}

func (p *Plugin) GetManifest() plugin.Manifest {
	return manifest
}

func (p *Plugin) Init(filename string, configPath string) {
	// load toml config
	_, err := toml.DecodeFile(configPath+"/"+filename+".toml", &config)
	if err != nil {
		panic(err)
	}
}

func (p *Plugin) IsConvertible(from router.Format, to router.Format) bool {
	if strings.ToLower(from.API) == "telegram-bot-api" && strings.ToLower(to.API) == "ubm-api" {
		if strings.ToLower(from.Method) == "update" && strings.ToLower(to.Method) == "receive" {
			if strings.ToLower(from.Protocol) == "http" {
				return true
			}
		}
		if strings.ToLower(from.Method) == "apiresponse" && strings.ToLower(to.Method) == "response" {
			if strings.ToLower(from.Protocol) == "http" {
				return true
			}
		}
	}
	if strings.ToLower(from.API) == "ubm-api" && strings.ToLower(to.API) == "telegram-bot-api" {
		if strings.ToLower(from.Method) == "send" && strings.ToLower(to.Method) == "apirequest" {
			if strings.ToLower(to.Protocol) == "http" {
				return true
			}
		}
	}
	return false
}

func (p *Plugin) Convert(packet router.Packet, to router.Format, ch router.Buffer) bool {
	from := packet.Head.Format
	if strings.ToLower(from.API) == "telegram-bot-api" && strings.ToLower(to.API) == "ubm-api" {
		if strings.ToLower(from.Method) == "update" && strings.ToLower(to.Method) == "receive" {
			switch strings.ToLower(from.Protocol) {
			case "http":
				return convertTgUpdateHttpToUbmReceive(packet, to, ch)
			}
		}
		if strings.ToLower(from.Method) == "apiresponse" && strings.ToLower(to.Method) == "response" {
			switch strings.ToLower(from.Protocol) {
			case "http":

			}
		}
	}
	if strings.ToLower(from.API) == "ubm-api" && strings.ToLower(to.API) == "telegram-bot-api" {
		if strings.ToLower(from.Method) == "send" && strings.ToLower(to.Method) == "apirequest" {
			switch strings.ToLower(from.Protocol) {
			case "http":

			}
		}
	}
	return false
}

var PluginInstance plugin.Converter = &Plugin{}