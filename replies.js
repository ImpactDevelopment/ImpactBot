const channels = require('./channels');

module.exports = [
  {
    pattern: /forge|installer/,
    message: '<#' + channels.faq + '> <#' + channels.upcoming + '>.',
    exclude: []
  },
  {
    pattern: /(web)?(site|page)|donate|become ?a? don(at)?or/,
    message: '[Click here](https://impactdevelopment.github.io).',
    exclude: []
  },
  {
    pattern: /issue|bug|crash|error|suggest(ion)?s?|feature|enhancement/,
    message: 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!',
    exclude: []
  },
  {
    pattern: /help|support|assistance/,
    message: 'Switch to the <#222120655594848256> channel!',
    exclude: [channels.help]
  },
  {
    pattern: /franky/,
    message: 'It does exactly what you think it does.',
    exclude: []
  },
  {
    pattern: /optifine/,
    message: 'Optifine can be installed with Impact 4.0 and up (except 4.2): [video](https://www.youtube.com/watch?v=o1LHq6L0ibk), [text](https://github.com/ImpactDevelopment/ImpactClient/blob/master/Optifine.md)',
    exclude: []
  },
  {
    pattern: /mediafire|direct (link|url|site|page)|adf\.?ly/,
    message: 'Are you looking for a direct link? Choose your version: [4.3](http://www.mediafire.com/file/9ujvsouklxoq5hj/Impact+4.3+for+Minecraft+1.12.2.zip).',
    exclude: []
  }
];
