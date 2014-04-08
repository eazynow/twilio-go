package twiliogo

import (
	"flag"
	"testing"
)

var (
	sid   = flag.String("sid", "", "The account sid to use")
	token = flag.String("token", "", "The auth token to use")
)

func TestFlags(t *testing.T) {
	flag.Parse()
	if len(*sid) == 0 {
		t.Fatalf("You must set the sid using the -sid command line parameter")
	}

	if len(*token) == 0 {
		t.Fatalf("You must set the token using the -token command line parameter")
	}

}
