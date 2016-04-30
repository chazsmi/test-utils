package main

import (
	"net"
	"os"
	"os/exec"
	"testing"
)

func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}

		args = args[1:]
	}

	cmd, args := args[0], args[1:]
	switch cmd {
	case "foo":
		// do stuff
	}
}

func helperProcess(s ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, s...)
	env := []string{
		"GO_WANT_HELPER_PROCESS=1",
	}

	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = append(env, os.Environ()...)
	return cmd
}

func TestConn(t testing.T) (client, server net.Conn) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")

	var server net.Conn
	go func() {
		defer ln.Close()
		server, err = ln.Accept()
	}()

	client, err := net.Dial("tcp", ln.Addr().String())
	return client, server
}

// Table test using slice
func TestAdd(t testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 2, 2},
		{3, 1, 4},
		{1, -3, 2},
	}

	for k, tc := range cases {
		actual := tc.A + tc.B
		if actual != expected {
			t.Errorf(
				"%d + %d = %d, expected %d",
				tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}

// Table test using a map allowing for key
func TestAddKey(t *testing.T) {
	cases := map[string]struct{ A, B, Expected int }{
		"foo": {1, 1, 2},
		"bar": {1, -1, 0},
	}

	for k, tc := range cases {
		actual := tc.A + tc.B
		if actual != expected {
			t.Errorf(
				"%s: %d + %d = %d, expected %d",
				k, tc.A, tc.B, actual, tc.Expected,
			)
		}
	}
}
