package main

import (
	"fmt"
	"os"
	"strconv"

	netfilter "github.com/guesslin/go-netfilter-queue"
)

var (
	QueueID     uint16 = 100
	MaxQueueBuf uint32 = 1024
)

func init() {
	if qid, err := strconv.ParseUint(os.Getenv("QID"), 10, 16); err == nil {
		QueueID = uint16(qid)
	}
	if qbuf, err := strconv.ParseUint(os.Getenv("QBUF"), 10, 32); err == nil {
		MaxQueueBuf = uint32(qbuf)
	}
}

func main() {
	var err error
	fmt.Println("===== start =====")

	nfq, err := netfilter.NewNFQueue(QueueID, MaxQueueBuf, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer nfq.Close()
	packets := nfq.GetPackets()

	for {
		select {
		case p := <-packets:
			fmt.Println(p.Packet)
			p.SetVerdict(netfilter.NF_ACCEPT)
		}
	}
}
