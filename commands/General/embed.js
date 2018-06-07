const { Command } = require('klasa');
const { MessageEmbed } = require('discord.js');

module.exports = class extends Command {
  constructor(...args) {
    super(...args, {
      name: 'embed',
      enabled: true,
      runIn: ['text', 'dm', 'group'],
      cooldown: 3,
      bucket: 1,
      aliases: [],
      permissionLevel: 6,
      requiredPermissions: ['SEND_MESSAGES', 'EMBED_LINKS'],
      requiredConfigs: [],
      description: 'Sends an embed message for you',
      quotedStringSupport: true,
      usage: '<title:str> [desc:str] [color:str]',
      usageDelim: ' ',
      extendedHelp: 'No extended help available.'
    });
  }
  
  async run(msg, [title, desc, color]) {
    const embed = new MessageEmbed().setTitle(title);
    if(desc) embed.setDescription(desc);
    if(color) embed.setColor(color.toUpperCase());
    msg.send(embed);
  }
}