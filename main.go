package gopheris

import (
	"fmt"

	"github.com/whitenoiseoss/gopheris/discord/intents"
	"github.com/whitenoiseoss/gopheris/internal"
)

func testRun() {
	bm := internal.NewBitmask()
	bm.AddFlags(intents.GUILDS, intents.GUILD_MESSAGE_REACTIONS)
	fmt.Println(bm.CheckFlag(intents.DIRECT_MESSAGES))
	fmt.Println(bm.CheckFlag((intents.GUILDS)))
	fmt.Println(bm.CheckFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS))
	fmt.Println(bm.CheckFlags(intents.GUILDS, intents.GUILD_MESSAGE_REACTIONS))
}
