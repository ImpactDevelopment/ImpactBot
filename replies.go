package main

import (
	"regexp"
)

type Reply struct {
	pattern string
	message string
	exclude []string
	regex   *regexp.Regexp
}

func init() {
	for i := range replies {
		r, err := regexp.Compile(replies[i].pattern)
		if err != nil {
			panic(err)
		}
		replies[i].regex = r
	}
}

var replies = []Reply{
	{
		pattern: `forge`,
		message: "Use the [installer](https://impactdevelopment.github.io) to install Forge (1.12.2 only)",
	},
	{
		pattern: `faq`,
		message: "[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)",
	},
	{
		pattern: `defender|virus|mcafee|norton`,
		message: "https://github.com/ImpactDevelopment/ImpactClient/wiki/Setup-FAQ#my-antivirus-says-the-installer-is-a-virus-is-it-a-virus\n\n[Direct download link after adfly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `tutorial`,
		message: "Tutorial videos for downloading and installing the client:\n[Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Mac OSX](https://www.youtube.com/watch?v=BBO0v4eq95k)\n[Linux](https://www.youtube.com/watch?v=k_29vgkPUbk)\n",
	},
	{
		pattern: `baritone\ssetting`,
		message: "[Baritone settings list and documentation](https://baritone.leijurv.com/baritone/api/Settings.html#field.detail)",
	},
	{
		pattern: `(take\sa?\s?)?screenshot`,
		message: "https://www.minecraft.net/en-us/article/screenshotting-guide",
	},
	{
		pattern: `use\sbaritone|baritone\susage|baritone\scommand|\.b|goal|goto|path`,
		message: "https://github.com/cabaletta/baritone/blob/master/USAGE.md",
	},
	{
		pattern: `installe?r?|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|ad\s?f\.?ly|(ad|u)\s?block|download|ERR_CONNECTION_ABORTED|evassmat|update|infect`,
		message: "[Direct download link after adfly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)",
	},
	{
		pattern: `lite\s*loader`,
		message: "[LiteLoader tutorial](https://github.com/ImpactDevelopment/ImpactClient/wiki/Adding-LiteLoader)",
	},
	{
		pattern: `(web\s?)?(site|page)`,
		message: "[Website](https://impactdevelopment.github.io)",
	},
	{
		pattern: `issue|bug|crash|error|suggest(ion)?s?|feature|enhancement`,
		message: "Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!",
	},
	{
		pattern: `help|support`,
		message: "Switch to the <#" + help + "> channel!",
		exclude: []string{help},
	},
	{
		pattern: `what(\sdoes|\sis|s|'s)?\s+franky`,
		message: "It does exactly what you think it does.",
	},
	{
		pattern: `opti\s*fine`,
		message: "Use the installer to add OptiFine to Impact!\n[Old tutorial](https://www.youtube.com/watch?v=o1LHq6L0ibk). [Text instructions](https://github.com/ImpactDevelopment/ImpactClient/wiki/Adding-OptiFine)",
	},
	{
		pattern: `macros?`,
		message: "You can edit macros in-game, click Impact Button then Macros.",
	},
	{
		pattern: `change(\s*logs?|s)`,
		message: "[Changelog](https://impactdevelopment.github.io/changelog)",
	},
	{
		pattern: `hack(s|ing|er|client)?`,
		message: "Please do not discuss hacks in this Discord.",
	},
	{
		pattern: `premium|donat|become\s*a?\s+don(at)?or`,
		message: "If you donate $5 or more, you will recieve early access to upcoming releases through nightly builds (**now including 1.14.4 builds!**), 1 premium mod (Ignite), a cape visible to other Impact users, a gold colored name in the Impact Discord Server, and access to #donator-help (with faster and nicer responses). Go on the [website](https://impactdevelopment.github.io/#donate) to donate.",
	},
	{
		pattern: `1\.14`,
		message: "Preview builds of 1.14.4 are available to donators. No ETA on full release.",
	},
	{
		pattern: `schematics?`,
		message: "[ONLY FOR 1.12 VERSION] Place the schematic file into `.minecraft/schematics` then type `.b build name`. Make sure the required blocks are in your hotbar!",
	},
	{
		pattern: `((crack|cracked) (launcher|account|game|minecraft))|(terramining|shiginima|(t(-|)launcher))`,
		message: "Impact does not support cracked launchers. You can attempt to use the unstable Forge version, but no further support will be provided.",
	},
}
