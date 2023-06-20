package config

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
)

func NewConfigIni(name string) {
	var serverCfg *ini.File
	serverCfg, err := ini.Load(name)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse %s: %v", name, err)
	}
	mapTo(serverCfg, "server", ServerSetting)
	fmt.Println(ServerSetting)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

var ServerSetting = &Server{}

type Server struct {
	HttpPort          int
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	WsPort            int
	WsPongWait        time.Duration
	WsReadBufferSize  int
	WsWriteBufferSize int
	WsWriteWait       time.Duration
	WsMaxMessageSize  int64
	IntranetIp        string
	MainRedisMode     string
	TeamRedisMode     string
	JwtSecret         string
}
