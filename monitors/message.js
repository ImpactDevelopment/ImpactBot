const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const replies = [
  [/4\.3|forge|installer/, '<#398628237753843732> <#451094829393510420>.', []],
  [/(web)?(site|page)|donate|become ?a? don(at)?or/, '[Click here](https://impactdevelopment.github.io).', []],
  [/issue|bug|crash|error|suggest(ion)?|feature|enhancement/, 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!', []],
  [/help|support|assistance/, 'Switch to the <#222120655594848256> channel!', ['222120655594848256']],
  [/franky/, 'It does exactly what you think it does.', []],
  [/optifine/, '[4.0, 4.1](https://www.youtube.com/watch?v=o1LHq6L0ibk), 4.2: not compatible', []]
];

module.exports = class extends Monitor {
  constructor(...args) {
    super(...args, {
      name: 'message',
      enabled: true,
      ignoreBots: true,
      ignoreSelf: true,
      ignoreOthers: false,
      ignoreWebhooks: true,
      ignoreEdits: true
    });
  }

  async run(msg) {
    const reply = replies
      .filter((e) => new RegExp('\\b' + e[0].toString().slice(1, -1) + '\\b', 'i').test(msg.content) && !e[2].includes(msg.channel.id))
      .reduce((a, e) => a + e[1] + ' ', '');

    if(reply) {
      const m = await msg.send(new MessageEmbed().setDescription(reply));
      setTimeout(m.delete.bind(m), 15000);
    }
  }
};