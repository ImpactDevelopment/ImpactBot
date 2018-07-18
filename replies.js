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
    message: '[4.0, 4.1](https://www.youtube.com/watch?v=o1LHq6L0ibk) | 4.2: not compatible | 4.3: <#' + channels.announcements + '>',
    exclude: []
  }
];