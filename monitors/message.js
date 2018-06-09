const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const replies = [
  [/\b(4\.3|forge|optifine|installer)\b/i, 'You\'ve mentioned an upcoming feature, <#398628237753843732> <#451094829393510420>.', []],
  [/\b((web)?(site|page)|donate|become ?a? don(at)?or)\b/i, 'Check out the [website](https://impactdevelopment.github.io/) :)', []],
  [/\b(issue|bug|crash|error|suggest(ion)?|feature|enhancement)\b/i, 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!', []],
  [/\b(help|support|assistance)\b/i, 'Switch to the <#222120655594848256> channel!', ['222120655594848256']],
  [/\b(franky)\b/i, 'It does exactly what you think it does.', []]
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
      .filter((e) => e[0].test(msg.content) && !e[2].includes(msg.channel.id))
      .reduce((a, e) => a + e[1] + ' ');

    if(reply) {
      const m = await msg.send(new MessageEmbed().setDescription(reply));
      setTimeout(m.delete.bind(m), 15000);
    }
  }
};