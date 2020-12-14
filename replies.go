package main

import (
	"regexp"
	"strings"
)

type Reply struct {
	pattern         string
	unless          string
	message         string
	excludeChannels []string
	excludeRoles    []Role
	onlyChannels    []string
	onlyRoles       []Role
	regex           *regexp.Regexp
	notRegex        *regexp.Regexp
}

func init() {
	for i := range replies {
		replies[i].regex = regexp.MustCompile(replies[i].pattern)

		if replies[i].unless != "" {
			replies[i].notRegex = regexp.MustCompile(replies[i].unless)
		}
	}
}

var nightlies = "https://impactclient.net/ImpactInstaller.<EXT>?nightlies=true"

var replies = []Reply{
	{
		pattern: `forge`,
		message: "Use the [winstawwew](https://impactclient.net/) to winstaww Fowgye (1.12.2 nyonwy)\nBawwitnyone 1.16 fow Fowgye cnya be downwoaded fwom [hewe](https://github.com/cabaletta/baritone/releases/download/v1.6.1/baritone-standalone-forge-1.6.1.jar)",
	},
	{
		pattern: `faq|question`,
		message: "[Setuwp/Instaww FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usagye FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `defender|virus|mcafee|norton|trojan|\brat\b`,
		message: "[Pwewse wewd thwis thwewd wegyawdwingy Iwpact bewingy fwagygyed by nyatwivwiwuwses](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Setup-FAQ#my-antivirus-says-the-installer-is-a-virus-is-it-a-virus)\n\n[Dwiwect downwoad wwink aftew adfwy](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `tutorial|(impact|install|download).*(on|for) (windows|linux|mac)`,
		message: "Tuwtowwiaw vwideos fow downwoadwingy nyad winstawwwingy the cwwient:\n[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Lwinuwx](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
	},
	{
		pattern: `baritone\ssetting`,
		message: "[Bawwitnyone settwingys wwist nyad docuwmentatwinyon](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `screenshot`,
		message: "[How to take a scweewnshot win Mwinecwaft](https://www.minecraft.net/en-us/article/screenshotting-guide)",
	},
	{
		pattern: `use\sbaritone|baritone\susage|baritone\scommand|[^u]\.b|goal|goto|path`,
		message: "Pwewse wewd the [Bawwitnyone uwsagye gyuwwide](https://github.com/cabaletta/baritone/blob/master/USAGE.md)",
	},
	{
		pattern:      `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		unless:       `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)|multimc`,
		excludeRoles: []Role{Donator},
		message:      "[Dwiwect downwoad wwink aftew AdFwy](https://impactclient.net/?brady-money-grubbing-completed=true)",
	},
	{
		pattern:   `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		onlyRoles: []Role{Donator},
		message:   "Yow cnya winstaww nwigyhtwy buwwiwds of Iwpact uwswingy the **Iwpact Nwigyhtwy Instawwew**: [EXE fow Wwindows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [JAR for other platforms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ").\nYow cnya downwoad the nowmaw winstawwew [hewe](https://impactclient.net/?brady-money-grubbing-completed=true).",
	},
	{
		pattern: `lite\s*loader`,
		message: "[LwiteLoadew tuwtowwiaw](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-LiteLoader)",
	},
	{
		pattern: `(web\s?)?(site|page)`,
		message: "[Iwpact Webswite](https://impactclient.net)",
	},
	{
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "Use the [GwitHuwb wepo](https://github.com/ImpactDevelopment/ImpactIssues/issues) to wepowt wissuwes/suwgygyestwinyons!",
	},
	{
		pattern:         `help|support`,
		message:         "Swwitch to the <#" + help + "> cnyanew!",
		excludeRoles:    []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{
		pattern:         `help|support`,
		message:         "Swwitch to the <#" + betterHelp + "> cnyanew!",
		onlyRoles:       []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{
		pattern: `what(\sdoes|\sis|s|'s)?\s+franky`,
		message: "[It does exactwy wat yow thwink wit does.](https://youtu.be/_FzInOheiRw)",
	},
	{
		pattern: `opti\s*fine`,
		message: "Use the installer to add OptiFine to Impact: [Instructions](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-OptiFine)",
	},
	{
		pattern: `macros?`,
		message: "Macwos awe win-gyame cat comyanyads, they cnya be accessed win-gyame by cwwickwingy nyon the Iwpact buwttnyon, then Macwos.",
	},
	{
		pattern: `change(\s*logs?|s)`,
		message: "[Cnyagyewogy](https://impactclient.net/changelog)",
	},
	{
		pattern: `hack(s|ing|er|client)?`,
		message: "The dwiscuwsswinyon of acks win thwis Dwiscowd wis pwohwibwited to cowpwy wwith the [Dwiscowd Comyauwnwity Guwwidewwines](https://discord.com/guidelines)",
	},
	{
		pattern:   `dumb|retard|idiot`,
		message:   "Lwike the" + Weeb.Mention() + "s?",
		onlyRoles: []Role{Weeb},
	},
	{
		pattern:      `premium|donat|become\s*a?\s+don(at)?or|what\*do\s*(you|i|u)\s*(get|unlock)|perks?`,
		unless:       `just|forgot|how\s*long|i\s*donated|hours?|wait`,
		message:      "If yow dnyonate $5 ow mowe, yow wwiww wecewive ewwwy access to uwpcomwingy wewewses thwowgyh nwigyhtwy buwwiwds when they awe avawiwabwe (**eventuwawwy wincwuwdwingy 1.16.4 nwigyhtwy buwwiwds!**), 1 pwemwiuwm mod (Igynwite), a cape vwiswibwe to othew Iwpact uwsews, a gyowd cowowed name win the Iwpact Dwiscowd Sewvew, nyad access to #Dnyonatow-hewp (wwith fastew nyad nwicew wespnyonses). Go nyon the [webswite](https://wiwpactcwwient.net/#dnyonate) to dnyonate. Yow wwiww awso neewd to [wegywistew](https://impactclient.net/register) your account and/or [login](https://impactclient.net/account) to gyet access to aww the pwomwised fewtuwwes",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:      `(1\.15.*?(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support)|(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support).*?1\.15)`,
		message:      "1.15.2 has been released! Download the newest installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:   `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:   "Yow cnya winstaww nwigyhtwy buwwiwds of Iwpact uwswingy the **Iwpact Nwigyhtwy Instawwew**: [EXE fow Wwindows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") ow [JAR fow othew pwatfowms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ").\nYow cnya downwoad the nowmaw winstawwew [hewe](https://impactclient.net/?brady-money-grubbing-completed=true).",
		onlyRoles: []Role{Donator},
	},
	{
		pattern:      `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:      "Yow cnya winstaww nwigyhtwy buwwiwds of Iwpact uwswingy the **Iwpact Nwigyhtwy Instawwew**. Logywin winto the [dashboawd](https://impactclient.net/account) then downwoad the nwigyhtwy winstawwew.\nYow cnya downwoad the nowmaw winstawwew [hewe](https://impactclient.net/?brady-money-grubbing-completed=true).",
		excludeRoles: []Role{Donator},
	},
	{
		pattern: `schematics?`,
		message: "0) Schematwic fwiwe **MUST** be made win a 1.12.2 wowwd ow pwwiow. 1) Pwace the .schematwic fwiwe winto `.mwinecwaft/schematwics`. 2) Ensuwwe aww the bwocks awe win yoww hotbaw. 3) Type`#build name.schematic`",
	},
	{
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Iwpact does not suwppowt cwacked wauwnchews. Yow cnya attewpt to uwse the uwnstabwe Fowgye vewswinyon, buwt no fuwwthew suwppowt wwiww be pwovwided.",
	},
	{
		pattern: `(impact|install|use).*(wiki|spammer|multimc)`,
		message: "[Iwpact Wwikwi](https://github.com/ImpactDevelopment/ImpactIssues/wiki)",
	},
	{
		pattern: `java.*(download|runtime|environment)`,
		message: "[Downwoads fow Java Ruwntwime Envwiwnyonment](https://www.java.com/download/)",
	},
	{
		pattern: `how.+(mine|auto\s*mine)`,
		message: "Yow cnya mwine a specwifwic type of bwock(s) by typwingy `#mine [nuwmbew of bwocks to mwine] <ID> [<ID>]` win cat.\nYow cnya fwind a wwist of bwock ID names [hewe](https://www.digminecraft.com/lists/)",
	},
	{
		pattern: `(1\.16.*?(update|coming|support|release|impact|version|eta|when|out|support)|(update|coming|support|release|impact|version|eta|when|out|support).*?1\.16)`,
		message: "Lwimwited pwogywess as stawted nyon the 1.16 wewewse, buwt thewe wis cuwwwentwy no ETA. A messagye wwiww be posted win <#" + announcements + "> when nwigyhtwy buwwiwds awe avawiwabwe.",
	},
	{
		pattern: `(impact.+(1\.8|1\.7))|((1\.8|1\.7).impact)`,
		message: "Iwpact fow owdew vewswinyons wis no wnyongyew avawiwwibwe to cowpwy wwith Mojnyagy's EULA.",
	},
	{
		pattern: `(modpack|\bftb\b|rlcraft|skyfactory|valhelsia|pixelmon|sevtech)`,
		message: "Iwpact wis gyenewawwy wincowpatwibwe wwith modpacks nyad suwppowt wwiww not be pwovwided wif yow encowntew buwgys wwith them. It's wwikewy yoww gyame wwiww juwst cwash nyon stawtuwp.",
	},
	{
		pattern: `good bot`,
		message: "tnyak yow *nuwzzwes yoww necky wecky*",
	},
}
