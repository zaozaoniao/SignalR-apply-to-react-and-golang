package services

import (
	"code.corp.elong.com/aos/Rosen-bridge/signalr"
	lru "github.com/hashicorp/golang-lru"
	log "github.com/sirupsen/logrus"
)

type AppHub struct {
	signalr.Hub
	Cache *lru.Cache
}

func (h *AppHub) SendPop(message string) {
	h.Clients().All().Send("realtime", message)
}

func (h *AppHub) RequestSyncTime() string {
	//h.Clients().All().Send("receive", time.Now().Format("2006/01/02 15:04:05"))
	if t, ok := h.Cache.Get("synctime"); ok {
		return t.(string)
	}
	return ""
}

func (h *AppHub) OnConnected(connectionID string) {
	// fmt.Printf("%s connected\n", connectionID)
	log.Infoln(connectionID, " connected\n")
}

func (h *AppHub) OnDisconnected(connectionID string) {
	log.Infoln(connectionID, " disconnected\n")
}
