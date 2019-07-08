const { Event } = require('klasa')
const Discord = require('discord.js')
module.exports = class extends Event {
  constructor (...args) {
    super(...args, { name: 'guildMemberAdd', enabled: true })
  }
  run (member) {
    let embed = new Discord.MessageEmbed()
      .setColor(3447003)
      .setTitle('Welcome to the Impact Discord!')
      .setDescription('In order to prevent spam and give you a chance to read the FAQs and rules, you will not be able to talk for ten minutes.\nIn the meantime, check the useful links below. Please do not DM a staff member while waiting. Try to resolve the problem using the FAQ, or the help channel when you can speak.')
      .addField('Setup/Install FAQ', '[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Setup-FAQ)', true)
      .addField('Usage FAQ', '[Click here!](https://github.com/ImpactDevelopment/ImpactClient/wiki/Usage-FAQ)', true)
      .addField('Rules', '[Click here!](https://discordapp.com/channels/208753003996512258/224684271913140224/306183650268020748)', true)
      .addField('Github Links', '[Impact](https://github.com/ImpactDevelopment/ImpactClient), [Installer](https://github.com/ImpactDevelopment/Installer/), [Baritone](https://github.com/cabaletta/baritone)', true)
      .addField('Downloading and installing the client', '[🔷 Windows](https://www.youtube.com/watch?v=QP6CN-1JYYE)\n[Linux 🐧](https://www.youtube.com/watch?v=k_29vgkPUbk)\nMacOS video will come soon. In the mean time, use the Linux video as the instructions are almost identical.')
      .setFooter('♿ Impact Client ♿')
    try {
      member.send({ embed })
    } catch (error) {
      console.log('A fucky wucky happened! (This user probably has PMs off?)')
      console.log(error)
    }
  }
}
