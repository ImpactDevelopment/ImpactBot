const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const channels = require('../channels');
const replies = require('../replies');

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
      .filter((e) => ((e.exclude && !e.exclude.includes(msg.channel.id)) || true) && 
              new RegExp('\\b(?:' + e.pattern.source + ')\\b', 'i').test(msg.content))
      .reduce((a, e) => a + e.message + ' ', '');

    if(reply && [channels.general, channels.help, channels.bot].includes(msg.channel.id)) {
      const m = await msg.send(new MessageEmbed().setDescription(reply));
      if(!msg.mentions.has(this.client.user)) {
        setTimeout(() => m.delete().catch(() => {}), 20000);
        let deletedUsingReaction = false;
        m.react('\uD83D\uDDD1'); // wastebasket emoji
        const collector = m.createReactionCollector((reaction, user) => reaction.emoji.name === '\uD83D\uDDD1' && user.id === msg.author.id, { time: 15000 });
        collector.on('collect', r => {
          m.delete().catch(() => {});
          deletedUsingReaction = true;
          collector.stop();
        });
        collector.on('end', () => !deletedUsingReaction ? m.delete().catch(() => {}) : null);
      }
    }
  }
};
