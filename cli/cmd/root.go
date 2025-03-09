package cmd

import (
	"flag"
	"log"

	"github.com/timetravel-1010/indexer/cli/cmd/program"
	"github.com/timetravel-1010/indexer/cli/internal/parser"
	"github.com/timetravel-1010/indexer/cli/internal/stdparser"
)

var useCustomImpl *bool

var (
	flags   = Flags{}
	indexer = program.Indexer{}
)

type Flags struct {
	directory *string
	zincURL   *string
	port      *string
	user      *string
	password  *string
	_index    *string
	_type     *string
}

func Execute() {
	re := program.HttpRequest{
		Creds: program.Credentials{
			User:     *flags.user,
			Password: *flags.password,
		},
		BaseURL: *flags.zincURL,
		Index:   *flags._index,
		Type:    *flags._type,
		Port:    *flags.port,
	}

	if !*useCustomImpl {
		log.Println("using std")
		indexer.Parser = stdparser.StdParser{}
	} else {
		log.Println("using custom")
		indexer.Parser = parser.Parser{}
	}
	indexer.Index(*flags.directory, re)
}

func init() {
	flags.directory = flag.String("dir", "enron_mail_20110402", "path to email directory")
	flags.zincURL = flag.String("zincurl", "localhost", "zincsearch host url")
	flags.port = flag.String("port", "4080", "zincsearch host port")
	flags.user = flag.String("user", "admin", "zincsearch username")
	flags.password = flag.String("password", "Complexpass#123", "zincsearch password")
	flags._index = flag.String("index", "enron", "index name")
	flags._type = flag.String("type", "_doc", "request payload type")
	useCustomImpl = flag.Bool("custom", false, "use custom implementation instead of std (net/mail) library")

	flag.Parse()
}
