package main

import (
	"flag"
)

type consumeDataMqtt struct {
}

func (consume consumeDataMqtt) Consume(rawdata []byte) error {
	deviceId, topic, payload, qos, consumType, err := ParseRawData(rawdata)
	if err != nil {
		return err
	}
	return PublishData(deviceId, topic, payload, qos, consumType)
}

func main() {
	host := flag.String("host", ":6377", "host of listen udp packeges")
	flag.Parse()
	var consumer ConsumeData
	consumer = new(consumeDataMqtt)
	UdpServer(*host, consumer);
}
