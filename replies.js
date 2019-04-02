const channels = require('./channels');

module.exports = [
  {
    pattern: /forge/,
    message: 'Very soon™ (<#' + channels.upcoming + '>).'
  },
  {
    pattern: /lite\s*loader/,
    message: '[LiteLoader tutorial](https://github.com/ImpactDevelopment/ImpactClient/blob/master/LiteLoader.md)',
  },
  {
    pattern: /(web)?(site|page)|donate|become ?a? don(at)?or/,
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
    pattern: /franky/,
    message: 'It does exactly what you think it does.'
  },
  {
    pattern: /opti\s*fine/,
    message: '[OptiFine tutorial (4.0+ but not 4.2)](https://www.youtube.com/watch?v=o1LHq6L0ibk), [text](https://github.com/ImpactDevelopment/ImpactClient/blob/master/Optifine.md)'
  },
  {
    pattern: /mediafire|dire(c|k)+to? (linko?|url|site|page)|adf\.?ly|(ad|u)block/,
    message: '[4.2](https://www.mediafire.com/file/ziqx4m44zkgj1ye/Impact+4.2+for+Minecraft+1.12.2.zip) | [4.3](https://www.mediafire.com/file/9ujvsouklxoq5hj/Impact+4.3+for+Minecraft+1.12.2.zip) | [4.4](https://www.mediafire.com/file/l7brss1f228so0p/Impact+4.4+for+Minecraft+1.12.2.zip) | [4.5](https://www.mediafire.com/file/a9srjpjfb4uppqj/Impact+4.5+for+Minecraft+1.12.2.zip). | [4.6](https://www.mediafire.com/file/ccjf2xs8veokkc9/Impact+4.6+for+Minecraft+1.13.2.zip)'
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
  }
];
