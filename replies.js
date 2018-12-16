const channels = require('./channels');

module.exports = [
  {
    pattern: /forge|installer/,
    message: 'Soon™ <#' + channels.faq + '> <#' + channels.upcoming + '>.'
  },
  {
    pattern: /liteloader/,
    message: '[Click here for an Impact + LiteLoader tutorial!](https://github.com/ImpactDevelopment/ImpactClient/blob/master/LiteLoader.md)',
  },
  {
    pattern: /(web)?(site|page)|donate|become ?a? don(at)?or/,
    message: '[Click here to open the website](https://impactdevelopment.github.io).'
  },
  {
    pattern: /issue|bug|crash|error|suggest(ion)?s?|feature|enhancement/,
    message: 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!'
  },
  {
    pattern: /help|support|assistance/,
    message: 'Switch to the <#' + channels.help + '> channel!',
    exclude: [channels.help]
  },
  {
    pattern: /franky/,
    message: 'It does exactly what you think it does.'
  },
  {
    pattern: /optifine/,
    message: 'Optifine can be installed with Impact 4.0 and up (except 4.2). Instructions: [video](https://www.youtube.com/watch?v=o1LHq6L0ibk), [text](https://github.com/ImpactDevelopment/ImpactClient/blob/master/Optifine.md)'
  },
  {
    pattern: /mediafire|die? ?re(c|k)+to? (linko?|url|site|page)|adf\.?ly|(ad|u)block/,
    message: 'Direct links: [4.2](http://www.mediafire.com/file/ziqx4m44zkgj1ye/Impact+4.2+for+Minecraft+1.12.2.zip) | [4.3](http://www.mediafire.com/file/9ujvsouklxoq5hj/Impact+4.3+for+Minecraft+1.12.2.zip) | [4.4](http://www.mediafire.com/file/l7brss1f228so0p/Impact+4.4+for+Minecraft+1.12.2.zip).'
  },
  {
    pattern: /macros/,
    message: '[Manually creating macros](https://github.com/ImpactDevelopment/ImpactClient/issues/153#issuecomment-399772723)'
  },
  {
    pattern: /changelog/,
    message: '[Changelog](https://impactdevelopment.github.io/changelog)'
  },
  {
    pattern: /damn/,
    message: 'Damn Daniel!'
  },
  {
    pattern: /op+a/,
    message: 'OPPA GANGNAM STYLE!'
  }
];
