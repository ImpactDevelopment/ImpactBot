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
    message: 'Optifine can be installed with Impact 4.0 and up (except 4.2): [video](https://www.youtube.com/watch?v=o1LHq6L0ibk), [text](https://github.com/ImpactDevelopment/ImpactClient/blob/master/Optifine.md)'
  },
  {
    pattern: /mediafire|direct (link|url|site|page)|adf\.?ly/,
    message: 'Direct links: [4.2](http://www.mediafire.com/file/ziqx4m44zkgj1ye/) | [4.3](http://www.mediafire.com/file/9ujvsouklxoq5hj/) | [4.4](http://www.mediafire.com/file/l7brss1f228so0p/).'
  },
  {
    pattern: /macros/,
    message: '[Manually creating macros](https://github.com/ImpactDevelopment/ImpactClient/issues/153#issuecomment-399772723)'
  },
  {
    pattern: /changelog/,
    message: '[Changelog](https://impactdevelopment.github.io/changelog)'
  }
];
