package poker

import "testing"

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	cli := &CLI{playerStore}
	cli.playPoker()
	if len(playerStore.winCalls) < 1 {
		t.Fatal("excepted a win call but didn't get one")
	}
}
