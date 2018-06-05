const { Client } = require('klasa');

Client.defaultPermissionLevels
	.add(9, (client, message) => client.options.owners.includes(message.author.id), { break: true })
	.add(10, (client, message) => client.options.owners.includes(message.author.id));

new Client({
  owners: process.env.OWNER_ID.split(' '),
  clientOptions: {
    fetchAllMembers: false,
    disableEveryone: true,
    disabledEvents: [/* Misc */
                     'TYPING_START','VOICE_STATE_UPDATE',
                     
                     /* Guild/Channel */
                     'CHANNEL_PINS_UPDATE','GUILD_SYNC',
                     
                     /* Message */
                     'MESSAGE_REACTION_REMOVE_ALL','MESSAGE_REACTION_REMOVE','MESSAGE_REACTION_ADD',
                    
                     /* Useless (some not used in bot accounts) */
                     'RELATIONSHIP_ADD','RELATIONSHIP_REMOVE','USER_NOTE_UPDATE',
                     'USER_SETTINGS_UPDATE']
  },
  prefix: 'i!',
  cmdEditing: true,
  readyMessage: (client) => `${client.user.tag}, Ready to serve ${client.guilds.size} guilds and ${client.users.size} users`
}).login(process.env.TOKEN);