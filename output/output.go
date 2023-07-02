package output

import (
	"strings"

	"github.com/cxfksword/httpcap/flags"
	"github.com/cxfksword/httpcap/protos"
)

type OutputPrinter interface {
	Print(data protos.Protos, srcAddr string, srcPort int, dstAddr string, dstPort int)
}

var printer []OutputPrinter

func Init() {
	printer = []OutputPrinter{
		&Console{
			showBody:       flags.Options.Body,
			showRaw:        flags.Options.Raw,
			showMoreDetail: flags.Options.MoreDetail,
		},
	}
}

func Print(data protos.Protos, srcAddr string, srcPort int, dstAddr string, dstPort int) {
	// filter by keyword
	if flags.Options.Keyword != "" && !strings.Contains(data.GetSearchKey(), flags.Options.Keyword) {
		return
	}

	for _, v := range printer {
		v.Print(data, srcAddr, srcPort, dstAddr, dstPort)
	}
}
