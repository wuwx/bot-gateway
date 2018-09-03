package router

import "time"

type Packet struct {
	Head Head
	Body []byte
}

type Head struct {
	UUID                  string
	From                  string
	To                    string
	ReplyToUUID           string
	AcknowlegeChannelUUID string
	Level                 int
	Format                Format
}

type Format struct {
	API      string
	Version  string
	Protocol string
}

type Buffer chan Packet

type Channel struct {
	UUID       string
	Buffer     *Buffer
	ExpireTime time.Time
}

type ProducerChannel struct {
	Channel
	AcknowlegeBuffer *Buffer
}

type ConsumerChannel struct {
	Channel
	Accept []RoutingRule
}

type RoutingRule struct {
	From    string
	To      string
	Level   int
	Formats []Format
}

func getExpireTime() time.Time {
	return time.Now().Local().Add(config.ChannelLifeTime)
}

func (ch Channel) renew() {
	ch.ExpireTime = getExpireTime()
}