const { Monitor } = require('klasa');
const { MessageEmbed } = require('discord.js');

const channels = require('../channels');
const replies = require('../replies');

const TIMEOUT = 20000;
const TRASH_EMOJI = '\uD83D\uDDD1';

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
    if(msg.member.roles.has('245682967546953738') && !msg.isMentioned(this.client.user)) return; // Don't spam out chat if a person with the Support role says the word and doesnt mention the bot
    const reply = replies
      .filter((e) => {
        if(e.exclude && e.exclude.includes(msg.channel.id)) return false;
        return new RegExp('\\b(?:' + e.pattern.source + ')\\b', 'i').test(msg.content);
      }).reduce((a, e) => a + e.message + ' ', '');

    if(reply && [channels.general, channels.help, channels.bot].includes(msg.channel.id)) {
      const m = await msg.send(new MessageEmbed().setDescription(reply));
      if(!msg.mentions.has(this.client.user)) {
        let deletedUsingReaction = false;
        m.react(TRASH_EMOJI);
        const collector = m.createReactionCollector(
          (reaction, user) => reaction.emoji.name === TRASH_EMOJI && user.id === msg.author.id, { time: TIMEOUT }
        );
        collector.on('collect', r => {
          m.delete().catch(() => {});
          deletedUsingReaction = true;
          collector.stop();
        });
        collector.on('end', () => !deletedUsingReaction && m.delete().catch(() => {}));
      }
    }
  }
};
