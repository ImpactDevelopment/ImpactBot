const channels = require('./channels');

module.exports = [
  {
    pattern: /forge/,
    message: 'Use the [installer](https://impactdevelopment.github.io) to install Forge (1.12.2 only)'
  },
  { 
    pattern: /faq/,
    message: '[Setup/Install FAQ](https://github.com/impactdevelopment/impactclient/wiki/Setup-FAQ)\n[Usage FAQ](https://github.com/impactdevelopment/impactclient/wiki/Usage-FAQ)'
  },
  {
    pattern: /(take\sa?\s?)?screenshot/,
    message: 'https://www.minecraft.net/en-us/article/screenshotting-guide'
  },
  {
    pattern: /use\sbaritone|baritone\susage/,
    message: 'https://github.com/cabaletta/baritone/blob/master/USAGE.md'
  },
  {
    pattern: /installer|mediafire|dire(c|k)+to?\s+(linko?|url|site|page)|adf\.?ly|(ad|u)\s?block/,
    message: '[Direct download link after adfly](https://impactdevelopment.github.io/?brady-money-grubbing-completed=true)'
  },
  {
    pattern: /lite\s*loader/,
    message: '[LiteLoader tutorial](https://github.com/ImpactDevelopment/ImpactClient/wiki/Adding-LiteLoader)',
  },
  {
    pattern: /(web\s?)?(site|page)/,
    message: '[Website](https://impactdevelopment.github.io)'
  },
  {
    pattern: /issue|bug|crash|error|suggest(ion)?s?|feature|enhancement/,
    message: 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!'
  },
  {
    pattern: /help|support/,
    message: 'Switch to the <#' + channels.help + '> channel!',
    exclude: [channels.help]
  },
  {
    pattern: /what(\sdoes|\sis|s|'s)?\s+franky/,
    message: 'It does exactly what you think it does.'
  },
  {
    pattern: /opti\s*fine/,
    message: 'Use the installer to add OptiFine to Impact!\n[Old tutorial](https://www.youtube.com/watch?v=o1LHq6L0ibk). [Text instructions](https://github.com/ImpactDevelopment/ImpactClient/wiki/Adding-OptiFine)'
  },
  {
    pattern: /directlinktomanualjsoninstallation/,
    message: '[4.2](https://www.mediafire.com/file/ziqx4m44zkgj1ye/Impact+4.2+for+Minecraft+1.12.2.zip) | [4.3](https://www.mediafire.com/file/9ujvsouklxoq5hj/Impact+4.3+for+Minecraft+1.12.2.zip) | [4.4](https://www.mediafire.com/file/l7brss1f228so0p/Impact+4.4+for+Minecraft+1.12.2.zip) | [4.5](https://www.mediafire.com/file/a9srjpjfb4uppqj/Impact+4.5+for+Minecraft+1.12.2.zip). | [4.6 1.12.2](http://www.mediafire.com/file/2ud3yig0i8i8p63/Impact+4.6+for+Minecraft+1.12.2.zip) | [4.6 1.13.2](https://www.mediafire.com/file/ccjf2xs8veokkc9/Impact+4.6+for+Minecraft+1.13.2.zip)'
  },
  {
    pattern: /macros?/,
    message: 'You can edit macros in-game, click Impact Button then Macros.'
  },
  {
    pattern: /change(\s*logs?|s)/,
    message: '[Changelog](https://impactdevelopment.github.io/changelog)'
  },
  {
    pattern: /hack(s|ing|er|client)?/,
    message: 'Please do not discuss hacks in this Discord.'
  },
  {
    pattern: /premium|donate|become\s*a?\s+don(at)?or/,
    message: 'If you donate $5 or more, you will recieve 1 premium mod (Ignite), a cape visible to other Impact users, a gold colored name in the Impact Discord Server, access to #donator-help (with faster and nicer responses), and access to closed betas of upcoming releases. Go on the website to donate.'
  },
  {
    pattern: /1\.14/,
    message: 'No release date for 1.14, a message will be in #announcements when development starts.'
  }
];
