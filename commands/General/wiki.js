const { Command } = require('klasa')
const { exec } = require('child_process')
const fs = require('fs')

function capitalizeFirstLetter (string) {
  return string.charAt(0).toUpperCase() + string.slice(1)
}

module.exports = class extends Command {
  constructor (...args) {
    super(...args, {
      name: 'wiki',
      enabled: true,
      runIn: ['text', 'dm'],
      cooldown: 3,
      bucket: 1,
      aliases: ['wiki', 'mod', 'module'],
      permissionLevel: 0,
      requiredPermissions: ['SEND_MESSAGES'],
      requiredConfigs: [],
      description: 'Get info from the Impact wiki.',
      quotedStringSupport: false,
      usage: '<category:str> <heading:str>',
      usageDelim: ';',
      extendedHelp: 'Get info from the Impact Wiki. All spellings are in American English. EG i!module Combat; AutoArmor or i!module Installation; Using the Installer'
    })
  }

  async run (msg, [category, heading]) {
    try {
      let data = fs.readFileSync('./ImpactClient.wiki/' + capitalizeFirstLetter(category) + '.md').toString()
      let res = data.split('## ' + capitalizeFirstLetter(heading)).pop().split('##')[0]
      if (res === '') {
        msg.channel.send('That module either doesn\'t exist, or you specified the wrong category.')
        msg.channel.send('The wiki is available here: https://github.com/ImpactDevelopment/ImpactClient/wiki')
        return
      }
      msg.channel.send('```fix\n' + res + '\n```')
    } catch (err) {
      msg.channel.send('That category does not exist!')
      msg.channel.send('Additional logging can be found below.')
      msg.channel.send('```\n' + err + '\n```')
      msg.channel.send('The wiki is available here: https://github.com/ImpactDevelopment/ImpactClient/wiki')
    }
  }
  async init () {
    exec('git clone https://github.com/ImpactDevelopment/ImpactClient.wiki.git', (err, stdout, stderr) => {
      if (err) {
        console.log('Cloning the wiki repo failed!')
        console.error(err)
      } else {
        console.log('Successfully cloned the wiki repository!')
      }
    })
  }
}
