const { Client, PermissionLevels } = require('klasa');
const { Permissions, Permissions: { FLAGS } } = require('discord.js');

Client.defaultPermissionLevels


module.exports = new Client({
  permissionLevels: new PermissionLevels()
      .add(0, () => true)
      .add(6, (client, message) => message.guild && message.member.permissions.has(FLAGS.MANAGE_GUILD), { fetch: true })
      .add(7, (client, message) => message.guild && message.member === message.guild.owner, { fetch: true })
      .add(9, (client, message) => client.options.owners.includes(message.author.id), { break: true })
      .add(10, (client, message) => client.options.owners.includes(message.author.id)),
  owners: process.env.OWNER_ID.split(' '),
  clientOptions: {
    fetchAllMembers: false,
    disableEveryone: true,
    disabledEvents: [/* Misc */ 'TYPING_START','VOICE_STATE_UPDATE',
                     /* Guild|Channel */ 'CHANNEL_PINS_UPDATE','GUILD_SYNC',
                     /* Message */ 'MESSAGE_REACTION_REMOVE_ALL','MESSAGE_REACTION_REMOVE','MESSAGE_REACTION_ADD',
                     /* Useless */ 'RELATIONSHIP_ADD','RELATIONSHIP_REMOVE','USER_NOTE_UPDATE',
                     'USER_SETTINGS_UPDATE']
  },
  prefix: 'i!',
  commandEditing: true,
  readyMessage: (client) => `${client.user.tag}, Ready to serve ${client.guilds.size} guilds and ${client.users.size} users`
}).login(process.env.TOKEN);