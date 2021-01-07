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
var installer = "https://impactclient.net/?brady-money-grubbing-completed=true"

var replies = []Reply{
	{
		pattern: `faq|question`,
		message: "[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `defender|virus|mcafee|norton|trojan|\brat\b`,
		message: "[FAQ: My antivirus says the installer is a virus! Is it a virus?](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Setup-FAQ#my-antivirus-says-the-installer-is-a-virus-is-it-a-virus)\n\n[Direct download link after AdFly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `tutorial|(impact|install|download).*(on|for) (windows|linux|mac)`,
		message: "Tutorial videos for downloading and installing the client:\n[Windows](https://www.youtube.com/watch?v=9IV_NC377pg)\n[macOS](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[GNU/Linux](https://www.youtube.com/watch?v=XPLvooJeQEI)\n",
	},
	{
		pattern: `baritone\s*setting`,
		message: "[Baritone settings list and documentation](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `screenshot`,
		message: "[How to take a screenshot in Minecraft](https://www.minecraft.net/en-us/article/screenshotting-guide)",
	},
	{
		pattern: `use\s*baritone|baritone\s*(usage|command)|[^u]\.b|goal|goto|path`,
		message: "Please read the [Baritone usage guide](https://github.com/cabaletta/baritone/blob/master/USAGE.md)",
	},
	{ // Info for non-donators about donating
		pattern: `premium|donat|become\s*a?\s+don(at)?or|what\s*do\s*(you|i|u)\s*(get|unlock)|perks?`,
		// Unless it would be a better fit for another reply (e.g. not asking about premium)
		unless: `(installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect)`,
		message: "If you donate $5 or more, you will receive early access to upcoming releases through nightly builds when they are available (**now including 1.16.4 nightly builds!**), " +
			"1 premium mod (Ignite), a cape visible to other Impact users, a gold colored name in the Impact Discord Server, and access to #donator-help (with faster and nicer responses). " +
			"Go on the [website](https://impactclient.net/#donate) to donate. You will also need to [register](https://impactclient.net/register) your account and/or " +
			"[login](https://impactclient.net/account) to get access to all the promised features",
		excludeRoles: []Role{Donator},
	},
	{ // Installer download for non-donators (not asking about premium)
		pattern: `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		// Unless it would be a better match for another reply (e.g. asking about premium, optifine or forge)
		unless:       `nightly|pre[- ]*release|beta|alpha|alfa|((download|get|where).*1[.]16)`,
		excludeRoles: []Role{Donator},
		message:      "Download the installer [here](" + installer + ") (direct download link without AdFly)",
	},
	{ // Installer download for donators
		pattern:   `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		onlyRoles: []Role{Donator},
		message: "You can download the normal installer [here](" + installer + ").\n" +
			"As a Donator, you can also install **nightly builds** of Impact: Download the Impact Nightly Installer [for Windows (.exe)](" + strings.Replace(nightlies, "<EXT>", "exe", 1) + ") or [for other platforms (.jar)](" + strings.Replace(nightlies, "<EXT>", "jar", 1) + ").\n" +
			"You can also access these links by logging into your [Impact Account dashboard](https://impactclient.net/account)",
	},
	{ // Install info for optifine
		pattern: `\b(install.*impact.*opti\s*fine)\b`,
		message: "Use the [installer](" + installer + ") to add OptiFine to Impact: [Instructions](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-OptiFine)",
	},
	{ // Install info for Forge
		pattern: `\b(install.*impact.*forge|support.*forge)\b`,
		message: "Use the [installer](https://impactclient.net/) to install Forge (1.12.2 only)\nStandalone Baritone supports Forge on various versions - Download from [GitHub](https://github.com/cabaletta/baritone/releases).",
	},
	{ // Install info for LiteLoader
		pattern: `lite\s*loader`,
		message: "[LiteLoader tutorial](https://github.com/ImpactDevelopment/ImpactIssues/wiki/Adding-LiteLoader)",
	},
	{ // Links to Impactt website
		pattern: `(web\s?)?(site|page)`,
		message: "[Impact Website](https://impactclient.net)",
	},
	{ // GitHub repo for issues
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactIssues/issues) to report issues/suggestions!",
	},
	{ // Non-donator help redirect
		pattern:         `help|support`,
		message:         "Switch to the <#" + help + "> channel!",
		excludeRoles:    []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{ // Donator help redirect
		pattern:         `help|support`,
		message:         "Switch to the <#" + betterHelp + "> channel!",
		onlyRoles:       []Role{Donator},
		excludeChannels: []string{help, betterHelp},
	},
	{ // What does Franky do????
		pattern: `what(\sdoes|\sis|s|'s)?\s+franky`,
		message: "[It does exactly what you think it does.](https://youtu.be/_FzInOheiRw)",
	},
	{ // Information on macros
		pattern: `macros?`,
		message: "Macros are in-game chat commands, they can be accessed in-game by clicking on the Impact button, then Macros.",
	},
	{ // Links to the Impact changelogs
		pattern: `change(\s*logs?|s)`,
		message: "[Changelog](https://impactclient.net/changelog)",
	},
	{ // Notice about "hacking"
		pattern: `hack(s|ing|er|client)?`,
		message: "**Impact is not a hacked client**, it is designed as a utility mod (e.g. for anarchy servers).\nSupport will not be provided to users who utilise Impact on servers that do not allow it.\nPlease also note that the discussion of \"hacks\" in this Discord server is prohibited to comply with the [Discord Community Guidelines](https://discord.com/guidelines)",
	},
	{ // Weeb moment
		pattern:   `dumb|retard|idiot`,
		message:   "Like the " + Weeb.Mention() + "s?",
		onlyRoles: []Role{Weeb},
	},
	{ // Info on using schematics
		pattern: `schematics?`,
		message: "Schematic file **MUST** be made in a 1.12.2 world or prior.\n1) Place the .schematic file into `.minecraft/schematics`.\n2) Ensure all the blocks are in your hotbar.\n3) Type `#build name.schematic`.",
	},
	{ // Info on using cracked launchers
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Impact does not support cracked launchers. You can attempt to use the unstable Forge version, but no further support will be provided.",
	},
	{ // Link to the Impact wiki
		pattern: `\b(impact\s*wiki|(setup|use)\s*spam(mer)?|faq)\b`,
		message: "[Impact Wiki](https://github.com/ImpactDevelopment/ImpactIssues/wiki)",
	},
	{ // Downloads for JRE
		pattern: `java.*(download|runtime|environment)`,
		message: "[Downloads for Java Runtime Environment](https://www.java.com/download/)",
	},
	{ // How to use Impact's automine function
		pattern: `how.+(mine|auto\s*mine)`,
		message: "You can mine a specific type of block(s) by typing `#mine [number of blocks to mine] <ID> [<ID>]` in chat.\nYou can find a list of block ID names [here](https://www.digminecraft.com/lists/)",
	},
	{ // Older versions of Impact
		pattern: `(impact.+(1\.8|1\.7))|((1\.8|1\.7).impact)`,
		message: "Impact for older versions is no longer availible to comply with Mojang's EULA.",
	},
	{ // Information on using Impact with modpacks
		pattern: `(modpack|\bftb\b|rlcraft|skyfactory|valhelsia|pixelmon|sevtech)`,
		message: "Impact is generally incompatible with modpacks and support will not be provided if you encounter bugs with them. It's likely your game will just crash on startup.",
	},
	{
		pattern: `good\s*bot`,
		message: "tnyak yow *nuwzzwes yoww necky wecky*",
	},
	{
		pattern: `((anti(-|\s*)(kb|knockback))|velocity)`,
		message: "**Velocity**, also known as **Anti-knockback**, is a module under \"Movement\" that prevents the player from taking knockback.",
	},
	{
		pattern: `(gui|r(-|\s)shift|module|(open|close|show|hide)\s*impact)`,
		message: "To open or close the Impact GUI, press the `rshift` key, located below `enter`.",
	},
}

