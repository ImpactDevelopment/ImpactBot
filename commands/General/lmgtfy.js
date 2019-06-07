    
const { Command } = require('klasa');
const { MessageEmbed } = require('discord.js');

module.exports = class extends Command {
  constructor(...args) {
    super(...args, {
      name: 'let me Google that for you :)',
      enabled: true,
      runIn: ['text', 'dm', 'group'],
      cooldown: 3,
      bucket: 1,
      aliases: [],
      permissionLevel: 0,
      requiredPermissions: ['SEND_MESSAGES', 'EMBED_LINKS'],
      requiredConfigs: [],
      description: 'Type a message and the bot will return a url',
      quotedStringSupport: true,
      usage: '<search:str>',
      usageDelim: ' ',
      extendedHelp: 'No extended help available.'
    });
  }
  
  async run(msg, [search]) {
    var searchString = search;
    var editstring = searchString.split(' ').join('+')
    var sendgay = 'http://lmgtfy.com/?q=' + editstring;
    msg.send(sendgay);
  }
}
