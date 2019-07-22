const { Command } = require('klasa')
const { MessageEmbed } = require('discord.js')
module.exports = class extends Command {
  constructor (...args) {
    super(...args, {
      name: 'lmgtfy',
      enabled: true,
      runIn: ['text'],
      cooldown: 2,
      bucket: 1,
      aliases: ['google', 'search'],
      permissionLevel: 0,
      requiredPermissions: ['SEND_MESSAGES'],
      requiredConfigs: [],
      description: 'For when people are too lazy to use google themselves.',
      quotedStringSupport: false,
      usage: '<query:str>',
      usageDelim: '',
      extendedHelp: ''
    })
  }

  async run (msg, [query]) {
    msg.delete()
    const embed = new MessageEmbed()
      .setDescription('https://lmgtfy.com/?q=' + query.replace(/ /g, '%20'))
    msg.channel.send({ embed })
  }
}
