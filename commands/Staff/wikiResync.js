const { Command } = require('klasa')
const { exec } = require('child_process')
const rimraf = require('rimraf')

module.exports = class extends Command {
  constructor (...args) {
    super(...args, {
      name: 'wikiResync',
      enabled: true,
      runIn: ['text'],
      cooldown: 30,
      bucket: 1,
      aliases: ['wikiupdate', 'wikipull', 'wikisync', 'updatewiki', 'syncwiki'],
      permissionLevel: 0,
      requiredPermissions: ['SEND_MESSAGES', 'CHANGE_NICKNAME'], // Lazy way to check for a staff role LMAO
      requiredConfigs: [],
      description: 'Sync the bot\'s local copy of the Wiki to the Wiki on Github.',
      quotedStringSupport: false,
      usage: '',
      usageDelim: '',
      extendedHelp: 'Update the bot\'s local copy of the wiki. Support Role+ only.'
    })
  }

  run (msg) {
    msg.channel.send('Working...')
    rimraf.sync('./ImpactClient.wiki')
    exec('git clone https://github.com/ImpactDevelopment/ImpactClient.wiki.git', (err, stdout, stderr) => {
      if (err) {
        msg.channel.send('Cloning the wiki repo failed! -- Please try to run this command again as soon as possible, as the bot no longer has a copy of the Wiki!')
        msg.channel.send('```\n' + err + '\n```')
      } else {
        msg.channel.send('Successfully cloned the wiki repository!')
        console.log('Updated the bot\'s local wiki!')
      }
    })
  }
}
