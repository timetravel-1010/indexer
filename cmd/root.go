package cmd

import (
	"flag"
	"log"

	"github.com/timetravel-1010/indexer/cmd/program"
	"github.com/timetravel-1010/indexer/internal/parser"
	"github.com/timetravel-1010/indexer/internal/stdparser"
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
	flags.directory = flag.String("dir", "enron_mail_20110402", "Path to the email directory.")
	flags.zincURL = flag.String("zincurl", "localhost", "url for zincsearch host.")
	flags.port = flag.String("port", "4080", "port for zincsearch host.")
	flags.user = flag.String("user", "admin", "Username of zincsearch client.")
	flags.password = flag.String("password", "Complexpass#123", "Password of zincsearch client.")
	flags._index = flag.String("index", "enron", "Name for the index.")
	flags._type = flag.String("type", "_doc", "Type of the post request payload.")
	useCustomImpl = flag.Bool("custom", false, "Use custom implementation instead of std library.")
	flag.Parse()

}
