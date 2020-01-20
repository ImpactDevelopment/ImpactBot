package main

import "time"

const (
	TIMEOUT = 30 * time.Second
	TRASH   = "🗑"

	announcements = "378645175947362320"
	general       = "208753003996512258"
	help          = "222120655594848256"
	distro        = "433342110461067266"
	bot           = "306182416329080833"
	betterGeneral = "617140506069303306"
	donatorHelp   = "583453983427788830"
	donatorInfo   = "613478149669388298"
	testing       = "617066818925756506"
)

var muteRoles = map[string]string{
	"":      "630800201015361566",
	help:    "230803433752363020",
	general: "352144990606196749",
	distro:  "624971877424693288",
	bot:     "640263788985188362",
}
