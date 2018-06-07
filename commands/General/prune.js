const { Command } = require('klasa');

module.exports = class extends Command {
  constructor(...args) {
    super(...args, {
      name: 'prune',
      enabled: true,
      runIn: ['text', 'dm', 'group'],
      cooldown: 3,
      bucket: 1,
      aliases: [],
      permissionLevel: 6,
      botPerms: ['SEND_MESSAGES', 'MANAGE_MESSAGES'],
      requiredConfigs: [],
      description: 'Bulk-removes messages from the current channel',
      quotedStringSupport: false,
      usage: '<amount:num> [user:user]',
      usageDelim: ' ',
      extendedHelp: 'No extended help available.'
    });
  }
  
  async run(msg, [amount, user]) {
    if(amount < 1 || amount > 98) return msg.send('Message amount range: 1-98');
    if(!msg.member.hasPermission('MANAGE_MESSAGES')) return msg.send('You don\'t have Manage Messages permission!');
    var msgs = await msg.channel.messages.fetch({limit: amount});
    if(user.id) msgs = msgs.filter(m => m.author.id == user.id);
    await msg.delete();
    msg.channel.bulkDelete(msgs);
  }
}