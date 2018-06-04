const { Command } = require('klasa');
const { MessageEmbed } = require('discord.js');

module.exports = class extends Command {
  constructor(...args) {
    super(...args, {
      name: 'embed',
      enabled: true,
      runIn: ['text', 'dm', 'group'],
      cooldown: 5,
      bucket: 1,
      aliases: [],
      permLevel: 6,
      botPerms: ['SEND_MESSAGES', 'EMBED_LINKS'],
      requiredConfigs: [],
      description: 'Sends an embed message for you',
      quotedStringSupport: true,
      usage: '<title:str> <desc:str> <color:str>',
      usageDelim: ' ',
      extendedHelp: 'No extended help available.'
    });
  }
  
  async run(msg, [title, desc, color]) {
    msg.send(new MessageEmbed().setTitle(title).setDescription(desc).setColor(color.toUpperCase()));
  }
}