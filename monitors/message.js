const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const list = [
  [/4\.3|forge(?!hax)|optifine|installer/, 'You\'ve mentioned an upcoming feature, check the faq and upcoming channels.', []],
  [/(web)?(site|page)|donate|become ?a? don(at)?or/, 'Check out the [website](https://impactdevelopment.github.io/) :)', []],
  [/issue|bug|crash|error|suggestion|feature|enhancement/, 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!', []],
  [/help|support|assistance/, 'Switch to the help channel!', ['222120655594848256']],
  [/franky/, 'It does exactly what you think it does.', []]
];

function buildStr(str, channelId) {
  let a = '';;
  list.forEach(e => {
    if(e[0].test(str) && !e[2].includes(channelId)) a += e[1] + ' ';
  });
  return a;
}

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
    const str = buildStr(msg.content, msg.channel.id);
    if(str) {
      const m = await msg.send(new MessageEmbed().setDescription(str));
      setTimeout(m.delete.bind(m), 15000);
    }
  }
};