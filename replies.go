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
		message: "Use the [installer](https://impactclient.net/) to install Forge (1.12.2 only)\nBaritone 1.16 for Forge can be downloaded from [here](https://github.com/cabaletta/baritone/releases/download/v1.6.1/baritone-standalone-forge-1.6.1.jar)",
	},
	{
		pattern: `faq|question`,
		message: "[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `defender|virus|mcafee|norton|trojan|\brat\b`,
		message: "**Impact is not a virus.** [Please read this thread regarding Impact being flagged by antiviruses](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Setup-FAQ#my-antivirus-says-the-installer-is-a-virus-is-it-a-virus)\n\n[Direct download link after adfly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `tutorial|(impact|install|download).*(on|for) (windows|linux|mac)`,
		message: "Tutorial videos for downloading and installing the client:\n[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Linux](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
	},
	{
		pattern: `baritone\s*setting`,
		message: "[Baritone settings list and documentation](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `(screenshot|screenie?)`,
		message: "[How to take a screenshot in Minecraft](https://www.minecraft.net/en-us/article/screenshotting-guide)",
	},
	{
		pattern: `use\sbaritone|baritone\susage|baritone\scommand|[^u]\.b|goal|goto|path`,
		message: "Please read the [Baritone usage guide](https://github.com/cabaletta/baritone/blob/master/USAGE.md) for information on using Baritone. To learn about Baritone's settings, please see [the Baritone settings documentation](https://baritone.leijurv.com/baritone/api/Settings.html).",
	},
	{
		pattern:      `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		unless:       `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)|multimc`,
		excludeRoles: []Role{Donator},
		message:      "[Direct download link after AdFly](https://impactclient.net/?brady-money-grubbing-completed=true)",
	},
	{
		pattern:   `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		onlyRoles: []Role{Donator},
		message:   "You can install nightly builds of Impact using the **Impact Nightly Installer**: [EXE for Windows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [JAR for other platforms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ").\nYou can download the normal installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
	},
	{
		pattern: `lite\s*loader`,
		message: "To use LiteLoader, [please follow this tutorial](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-LiteLoader).",
	},
	{
		pattern: `(web\s?)?(site|page)`,
		message: "[Impact Website](https://impactclient.net)",
	},
	{
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "If you have encountered a bug whilst using Impact that cannot be answered in the help channels, or have a suggestion for a feature, use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactIssues/issues) to report issues/suggestions!",
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
		message: "Use the installer to add OptiFine to Impact: [Instructions](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-OptiFine)",
	},
	{
		pattern: `macro?s?`,
		message: "Macros are in-game chat commands, they can be accessed in-game by opening the pause menu, clicking on the `Impact` button, then `Macros`.",
	},
	{
		pattern: `change(\s*logs?|s)`,
		message: "[Changelog](https://impactclient.net/changelog)",
	},
	{
		pattern: `hack(s|ing|er|client)?`,
		message: "The discussion of hacks in this Discord is prohibited to comply with the [Discord Community Guidelines](https://discord.com/guidelines)",
	},
	{
		pattern:   `dumb|retard|idiot`,
		message:   "Like the " + Weeb.Mention() + "s?",
		onlyRoles: []Role{Weeb},
	},
	{
		pattern:      `premium|donat|become\s*a?\s+don(at)?or|what\*do\s*(you|i|u)\s*(get|unlock)|perks?`,
		unless:       `just|forgot|how\s*long|i\s*donated|hours?|wait`,
		message:      "If you donate $5 or more, you will receive early access to upcoming releases through nightly builds when they are available (**eventually including 1.16.4 nightly builds!**), 1 premium mod (Ignite), a cape visible to other Impact users, a gold colored name in the Impact Discord Server, and access to #Donator-help (with faster and nicer responses). Go on the [website](https://impactclient.net/#donate) to donate. You will also need to [register](https://impactclient.net/register) your account and/or [login](https://impactclient.net/account) to get access to all the aformentioned features",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:      `(1\.15.*?(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support)|(fucking|get|where|need|asap|update|coming|support|release|impact|version|eta|when|out|support).*?1\.15)`,
		message:      "The current version of Impact is Minecraft release 1.15.2. Download the installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
		excludeRoles: []Role{Donator},
	},
	{
		pattern:   `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:   "You can install nightly builds of Impact using the **Impact Nightly Installer**: [EXE for Windows](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [JAR for other platforms](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ").\nYou can download the normal installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
		onlyRoles: []Role{Donator},
	},
	{
		pattern:      `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]15)`,
		message:      "You can install nightly builds of Impact using the **Impact Nightly Installer**. Login into the [dashboard](https://impactclient.net/account) then download the nightly installer.\nYou can download the normal installer [here](https://impactclient.net/?brady-money-grubbing-completed=true).",
		excludeRoles: []Role{Donator},
	},
	{
		pattern: `schematics?`,
		message: "To use a schematic file, please carry out the following:\n0) Schematic file **MUST** be made in a 1.12.2 world or prior. 1) Place the .schematic file into `.minecraft/schematics`. 2) Ensure all the blocks are in your hotbar. 3) Type `#build name.schematic`",
	},
	{
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Impact does not support cracked launchers. You can attempt to use the unstable Forge version, but no further support will be provided.",
	},
	{
		pattern: `(impact|install|use).*(wiki|spammer|multimc)`,
		message: "For information on how to install and use Impact, see the [Impact Wiki](https://github.com/ImpactDevelopment/ImpactIssues/wiki)",
	},
	{
		pattern: `java.*(download|runtime|environment)`,
		message: "Impact requires Java Runtime Enviroment 8 to run.\n[Downloads for Java Runtime Environment](https://www.java.com/download/)",
	},
	{
		pattern: `how.+(mine|auto\s*mine)`,
		message: "You can mine a specific type of block(s) by typing `#mine [number of blocks to mine] <ID> [<ID>]` in chat.\nYou can find a list of block ID names [here](https://www.digminecraft.com/lists/)",
	},
	{
		pattern: `(1\.16.*?(update|coming|support|release|impact|version|eta|when|out|support)|(update|coming|support|release|impact|version|eta|when|out|support).*?1\.16)`,
		message: "Limited progress has started on the 1.16 release, but there is currently no ETA. A message will be posted in <#" + announcements + "> when nightly builds are available.",
	},
	{
		pattern: `(impact.+(1\.8|1\.7))|((1\.8|1\.7).impact)`,
		message: "Impact for older versions of Minecraft (pre 1.8) is no longer availible to comply with Mojang's EULA.",
	},
	{
		pattern: `(modpack|\bftb\b|rlcraft|skyfactory|valhelsia|pixelmon|sevtech)`,
		message: "Impact is generally incompatible with modpacks and support will not be provided if you encounter bugs with them. It's likely your game will just crash on startup.",
	},
	{
		pattern: `(good\s*(bot|human))`,
		message: "thank you \*nuzzles*",
	},
}
