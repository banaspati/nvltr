package command

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/redite/tlbot"
)

func init() {
	register(cmdHelp)
}

var cmdHelp = &Command{
	Name:      "help",
	ShortLine: "how to use this shit?",
	Run:       runHelp,
	Hidden:    true,
}

func runHelp(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	opts := &tlbot.SendOptions{
		ReplyTo:   msg.ID,
		ParseMode: tlbot.ModeNone,
	}
	b.SendMessage(msg.Chat.ID, help(), opts)
}

type byName []*Command

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func help() string {
	var buf bytes.Buffer

	var cmds []*Command
	for _, cmd := range commands {
		cmds = append(cmds, cmd)
	}

	sort.Sort(byName(cmds))

	buf.WriteString("Available Command:\n\n")
	for _, cmd := range cmds {
		// do not include hidden commands
		if cmd.Hidden {
			continue
		}
		buf.WriteString(fmt.Sprintf("/%v - %v\n", cmd.Name, cmd.ShortLine))
	}

	return buf.String()
}
