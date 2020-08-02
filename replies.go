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
		message: "Use the [installer](https://impactclient.net/) to install Forge (1.12.2 only)",
	},
	{
		pattern: `faq`,
		message: "[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `defender|virus|mcafee|norton`,
		message: "https://github.com/ImpactDevelopment/ImpactIssues/wiki/Setup-FAQ#my-antivirus-says-the-installer-is-a-virus-is-it-a-virus\n\n[Direct download link after adfly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `tutorial`,
		message: "Tutorial videos for downloading and installing the client:\n[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Linux](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
	},
	{
		pattern: `baritone\ssetting`,
		message: "[Baritone settings list and documentation](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `(take\sa?\s?)?screenshot`,
		message: "[How to take a screenshot in Minecraft](https://www.minecraft.net/en-us/article/screenshotting-guide)",
	},
	{
		pattern: `use\sbaritone|baritone\susage|baritone\scommand|\.b|goal|goto|path`,
		message: "[Baritone usage guide](https://github.com/cabaletta/baritone/blob/master/USAGE.md)",
	},
	{
		pattern:      `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		unless:       `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)|multimc`,
		excludeRoles: []Role{Donator},
		message:      "[Direct download link after adfly](https://impactclient.net/?brady-money-grubbing-completed=true)",
	},
	{
		pattern:   `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		onlyRoles: []Role{Donator},
		message:   "You can install nightly builds of Impact using the **Impact Nightly Installer**: [EXE for Windows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [JAR for other platforms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ")",
	},
	{
		pattern: `lite\s*loader`,
		message: "[LiteLoader tutorial](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-LiteLoader)",
	},
	{
		pattern: `(web\s?)?(site|page)`,
		message: "[Impact Website](https://impactclient.net)",
	},
	{
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactIssues/issues) to report issues/suggestions!",
	},
	{
		pattern:         `help|support`,
		message:         "Switch to the <#" + help + "> channel!",
		excludeRoles:    []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{
		pattern:         `help|support`,
		message:         "Switch to the <#" + betterHelp + "> channel!",
		onlyRoles:       []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{
		pattern: `what(\sdoes|\sis|s|'s)?\s+franky`,
		message: "[It does exactly what you think it does.](https://youtu.be/_FzInOheiRw)",
	},
	{
		pattern: `opti\s*fine`,
		message: "Use the installer to add OptiFine to Impact! [Text instructions](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-OptiFine)",
	},
	{
		pattern: `macros?`,
		message: "Macros are in-game chat commands, they can be accessed in-game by clicking on the Impact button then Macros.",
	},
	{
		pattern: `change(\s*logs?|s)`,
		message: "[Changelog](https://impactclient.net/changelog)",
	},
	{
		pattern: `hack(s|ing|er|client)?`,
		message: "Please do not discuss hacks in this Discord.",
	},
	{
		pattern:   `dumb|retard|idiot`,
		message:   "Like the <@&612744883467190275>s?",
		onlyRoles: []Role{{"612744883467190275", "fucking weeb"}},
	},
	{
		pattern:      `premium|donat|become\s*a?\s+don(at)?or|what\*do\s*(you|i|u)\s*(get|unlock)|perks?`,
		unless:       `just|forgot|how\s*long|i\s*donated|hours?|wait`,
		message:      "If you donate $5 or more, you will receive early access to upcoming releases through nightly builds (**now including 1.15.2 builds!**), 1 premium mod (Ignite), a cape visible to other Impact users, a gold colored name in the Impact Discord Server, and access to #Donator-help (with faster and nicer responses). Go on the [website](https://impactclient.net/#donate) to donate. You will also need to [register](https://impactclient.net/register) your account and/or [login](https://impactclient.net/account) to get access to all the promised features",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:      `(1\.15.*?(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support)|(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support).*?1\.15)`,
		message:      "1.15.2 support is now out! Download the newest installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:   `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:   "You can install nightly builds of Impact using the **Impact Nightly Installer**: [EXE for Windows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [JAR for other platforms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ")",
		onlyRoles: []Role{Donator},
	},
	{
		pattern:      `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:      "You can install nightly builds of Impact using the **Impact Nightly Installer**. Login into the [dashboard](https://impactclient.net/account) then download the nightly installer.",
		excludeRoles: []Role{Donator},
	},
	{
		pattern: `schematics?`,
		message: "0) Schematic file **MUST** be made in a 1.12.2 world or prior. 1) Place the .schematic file into `.minecraft/schematics`. 2) Ensure all the blocks are in your hotbar. 3) Type `#build name.schematic`",
	},
	{
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Impact does not support cracked launchers. You can attempt to use the unstable Forge version, but no further support will be provided.",
	},
	{
		pattern: `(impact|install|use).*(wiki|spammer|multimc)`,
		message: "Impact Wiki: https://github.com/ImpactDevelopment/ImpactIssues/wiki",
	},
	{
		pattern: `java.*(download|runtime|environment)`,
		message: "Java download: https://www.java.com/download/",
	},
	{
		pattern: `how.+(mine|auto\s*mine)`,
		message: "You can mine a specific type of block(s) by typing `#mine <ID> [<ID>]` in chat.\nYou can find a list of block ID names [here](https://www.digminecraft.com/lists/)",
	},
	{
		pattern: `(1\.16.*?(update|coming|support|release|impact|version|eta|when|out|support)|(update|coming|support|release|impact|version|eta|when|out|support).*?1\.16)`,
		message: "No ETA on 1.16 Impact release, a message will be posted in <#" + announcements + "> when development starts & nightly builds.",
	},
	{
		pattern: `(impact.+(1\.8|1\.7))|((1\.8|1\.7).impact)`,
		message: "Impact for older versions is no longer availible to comply with Mojang's EULA.",
	},
}
