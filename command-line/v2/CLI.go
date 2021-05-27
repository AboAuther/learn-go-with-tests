package poker

type CLI struct {
	playStore PlayerStore
}

func (cli *CLI) playPoker() {
	cli.playStore.RecordWin("lrg")
}
