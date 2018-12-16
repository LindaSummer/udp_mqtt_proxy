package main

import (
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"errors"
)

var (
	clients = make(map[string] mqtt.Client)
)

type DataType int

const (
	Publish = iota
	Subscribe
	Will
)

//ParseRawData
func ParseRawData(rawdata []byte) (string, string, []byte, int32, DataType, error) {
	return "", "", nil, 0, 0, nil
}

//PublishData
func PublishData(deviceId string, topic string, payload []byte, qos int32, dataType DataType) error {
	val, ok := clients[deviceId]
	if ok {
		token:= val.Publish(topic, byte(qos), false, payload);
		if token.WaitTimeout(10 * time.Second) {
			return nil;
		} else {
			return errors.New("public failed")
		}
	}
	if dataType == Will {
		options := mqtt.NewClientOptions();
		options.SetAutoReconnect(true)
		options.SetWill(topic, string(payload), byte(qos), false)
		clients[deviceId] = mqtt.NewClient(options)
	}
	return nil
}