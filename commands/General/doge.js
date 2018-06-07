const snekfetch = require('snekfetch');
const { Command } = require('klasa');

module.exports = class extends Command {
  constructor(...args) {
    super(...args, {
      name: 'doge',
      enabled: true,
      runIn: ['text', 'dm', 'group'],
      cooldown: 3,
      bucket: 1,
      aliases: ['shiba', 'shibe'],
      permissionLevel: 0,
      requiredPermissions: ['SEND_MESSAGES', 'ATTACH_FILES'],
      requiredConfigs: [],
      description: 'Sends a random doge image/gif in the channel 🐶',
      quotedStringSupport: false,
      usage: '',
      usageDelim: ' ',
      extendedHelp: 'No extended help available.'
    });
  }
  
  async run(msg, [...params]) {
    var res = await snekfetch.get('http://shibe.online/api/shibes?count=1&urls=true');
		msg.channel.send({files: [res.body[0]]});
  }
}