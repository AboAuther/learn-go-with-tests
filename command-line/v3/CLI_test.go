package poker

import (
	poker "github.com/quii/learn-go-with-tests/command-line/v3"
	"io"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record lrg win from user input", func(t *testing.T) {
		in := strings.NewReader("lrg wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "lrg")
	})
	t.Run("record chj win from user input", func(t *testing.T) {
		in := strings.NewReader("chj wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "chj")
	})
	t.Run("do not read beyond the first newline", func(t *testing.T) {
		in := failOnEndReader{
			t,
			strings.NewReader("lrg wins\n hello there"),
		}
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()
	})
}

type failOnEndReader struct {
	t   *testing.T
	rdr io.Reader
}

func (m failOnEndReader) Read(p []byte) (n int, err error) {
	n, err = m.rdr.Read(p)
	if n == 0 || err == io.EOF {
		m.t.Fatal("read to the end when you shouldn't have")
	}
	return n, err
}
