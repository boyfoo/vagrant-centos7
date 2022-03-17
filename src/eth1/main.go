package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func main() {

	// 需要切换到root 否则可能是空数组
	devs, err := pcap.FindAllDevs()
	if err != nil {
		return
	}
	for _, dev := range devs {
		fmt.Println(dev.Name)
	}

	handler, err := pcap.OpenLive("eth1", 1024, false, 3*time.Second)
	if err != nil {
		log.Fatalln(err)
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for packets := range source.Packets() {
		//fmt.Println(packets.String())
		// 第四层
		if layer4 := packets.TransportLayer(); layer4 != nil {
			// 是否是tcp (因为可能是udp)，并且是8080端口的数据
			if tcplayer, ok := layer4.(*layers.TCP); ok {

				// DstPort 目标 SrcPort源端口 (例子: 访问 客户端61132 -> 9090   回复时 9090 ->61132)
				if tcplayer.DstPort == 9090 {

					fmt.Printf("客户端%d-->服务端%d, SYN=%v,ACK=%v,Payload length=%d \n",
						tcplayer.SrcPort,
						tcplayer.DstPort,
						tcplayer.SYN,
						tcplayer.ACK,
						len(tcplayer.Payload),
					)

					if len(tcplayer.Payload) != 0 {
						fmt.Println("==============接受内容==============")
						fmt.Println(string(tcplayer.Payload))
					}
				}

				if tcplayer.SrcPort == 9090 {

					fmt.Printf("服务端%d-->客户端%d, SYN=%v,ACK=%v,Payload length=%d \n",
						tcplayer.SrcPort,
						tcplayer.DstPort,
						tcplayer.SYN,
						tcplayer.ACK,
						len(tcplayer.Payload),
					)

					if len(tcplayer.Payload) != 0 {
						fmt.Println("==============返回内容==============")
						fmt.Println(string(tcplayer.Payload))
					}
				}
			}
		}
	}
}
