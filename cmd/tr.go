package cmd

import (
	"fmt"

	"net"
	"os"
	"time"

	logging "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var (
	host = "www.baidu.com"
)
var TRCmd = &cobra.Command{
	Use: "traceroute",
	Run: func(cmd *cobra.Command, args []string) {

		ips, err := net.LookupIP(host)
		if err != nil {
			logging.WithError(err).Fatalf("failed to lookup IP %s", host)
		}
		var dst net.IPAddr
		for _, ip := range ips {
			if ip.To4() != nil {
				dst.IP = ip
				logging.Infof("using %v for tracing an IP packet route to %s\n", dst.IP, host)
				break
			}
		}
		if dst.IP == nil {
			logging.Fatal("no A record found")
		}

		c, err := net.ListenPacket("ip4:1", "0.0.0.0") // ICMP for IPv4
		if err != nil {
			logging.WithError(err).Fatal("failed to list packet")
		}
		defer c.Close()
		p := ipv4.NewPacketConn(c)

		if err := p.SetControlMessage(ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
			logging.WithError(err).Fatal("failed to set control message")
		}
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Data: []byte("Hello World"),
			},
		}

		rb := make([]byte, 1500)
		for i := 1; i <= 64; i++ { // up to 64 hops
			wm.Body.(*icmp.Echo).Seq = i
			wb, err := wm.Marshal(nil)
			if err != nil {
				logging.WithError(err).Fatal("failed to marshal data")
			}
			if err := p.SetTTL(i); err != nil {
				logging.WithError(err).Fatal("failed to set TTL")
			}

			// In the real world usually there are several
			// multiple traffic-engineered paths for each hop.
			// You may need to probe a few times to each hop.
			begin := time.Now()
			if _, err := p.WriteTo(wb, nil, &dst); err != nil {
				logging.WithError(err).Fatalf("failed to write")
			}
			if err := p.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
				logging.WithError(err).Fatalf("failed to read deadline")
			}
			n, cm, peer, err := p.ReadFrom(rb)
			if err != nil {
				if err, ok := err.(net.Error); ok && err.Timeout() {
					fmt.Printf("%v\t*\n", i)
					continue
				}
				logging.WithError(err).Fatal("failed to read")
			}
			rm, err := icmp.ParseMessage(1, rb[:n])
			if err != nil {
				logging.WithError(err).Fatalf("failed to parse icmp message")
			}
			rtt := time.Since(begin)

			// In the real world you need to determine whether the
			// received message is yours using ControlMessage.Src,
			// ControlMessage.Dst, icmp.Echo.ID and icmp.Echo.Seq.
			switch rm.Type {
			case ipv4.ICMPTypeTimeExceeded:
				names, _ := net.LookupAddr(peer.String())
				fmt.Printf("%d\t%v %+v %v\n\t%+v\n", i, peer, names, rtt, cm)
			case ipv4.ICMPTypeEchoReply:
				names, _ := net.LookupAddr(peer.String())
				fmt.Printf("%d\t%v %+v %v\n\t%+v\n", i, peer, names, rtt, cm)
				return
			default:
				logging.Errorf("unknown ICMP message: %+v\n", rm)
			}
		}
	},
}
