const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const list = [
  [/4\.3|forge|optifine|installer/, 'You\'ve mentioned an upcoming feature, please wait for it and check the faq or upcoming chann'],
  [/(web)?(site|page)|donate|become ?a? don(at)?or/, 'Check out the [website](https://impactdevelopment.github.io/) :)'],
  [/issue|bug|crash|error|suggestion|feature|enhancement/, 'Use the [GitHub repo](https://github.com/ImpactDevelopment/ImpactClient/issues) to report issues/suggestions!'],
  [/help|support|assistance/, 'Switch to the help channel!']
];

function buildStr(str) {
  let a = '';;
  list.forEach(e => {
    if(e[0].test(str)) a += e[1] + ' ';
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

  run(msg) {
    const str = buildStr(msg.content);
    if(str) msg.send(new MessageEmbed().setDescription(str));
  }
};