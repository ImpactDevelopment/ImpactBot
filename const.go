package main

//noinspection GoUnusedConst
const (
	impactServer     = "208753003996512258"
	brady            = "205718273696858113"
	prettyembedcolor = 0x3498DB

	trash = "🗑"
	check = "✅"

	announcements = "378645175947362320"
	general       = "208753003996512258"
	media         = "514880901629607947"
	help          = "222120655594848256"
	distro        = "433342110461067266"
	politics      = "745728642587295744"
	bot           = "306182416329080833"
	donatorInfo   = "613478149669388298"
	betterGeneral = "617140506069303306"
	betterHelp    = "583453983427788830"
	na            = "616409814779691114"
	oldguys       = "293796603146272768"
	development   = "280478531346104321"
	impactBotLog  = "617549691730526209"
	testing       = "617066818925756506"
)

var (
	HeadDev   = Role{"209817890713632768", "Head Developer"}
	Developer = Role{"221655083748687873", "Developer"}
	SeniorMod = Role{"663065117738663938", "Senior Moderator"}
	Moderator = Role{"210377982731223040", "Moderator"}
	Support   = Role{"245682967546953738", "Support"}
	Donator   = Role{"210114021641289728", "Donator"}
	Verified  = Role{"671048798654562354", "Verified"}
	InVoice   = Role{"677329885680762904", "In Voice"}
	Stupid    = Role{"743903534160019476", "Stupid"}
	Weeb      = Role{"612744883467190275", "weeb"}
	Givaway   = Role{"698619050833477633", "Givaway Ping"}
)

type Role struct {
	// Discord id
	ID   string
	Name string
}

func RolesToIDs(roles []Role) []string {
	var ids []string
	for _, role := range roles {
		ids = append(ids, role.ID)
	}
	return ids
}

var muteRoles = map[string]string{
	"":            "630800201015361566",
	general:       "352144990606196749",
	media:         "669632558551793666",
	help:          "230803433752363020",
	distro:        "624971877424693288",
	politics:      "745729048688328865",
	bot:           "640263788985188362",
	betterGeneral: "669633342525800471",
	betterHelp:    "669633463242194955",
	na:            "669632725644214283",
	oldguys:       "669632828371107881",
	development:   "669632988686057472",
}

func (r Role) Mention() string {
	return "<@&" + r.ID + ">"
}
