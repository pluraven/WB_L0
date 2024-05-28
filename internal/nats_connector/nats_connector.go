package nats_connector

import (
	"github.com/nats-io/stan.go"
)

func Run(ClusterID string, ClientID string, ChannelName string, URL string, messageHandler stan.MsgHandler) (stan.Conn, error) {
	connection, err := stan.Connect(ClusterID, ClientID, stan.NatsURL(URL))
	if err != nil {
		return nil, err
	}
	_, err = connection.Subscribe(ChannelName, messageHandler, stan.DurableName(ClientID))
	if err != nil {
		return nil, err
	}
	return connection, nil
}
