package rtspserver

import "github.com/angelodlfrtr/dorsvr/livemedia"

type StreamServerState struct {
	subsession  livemedia.IServerMediaSubsession
	streamToken *livemedia.StreamState
}
