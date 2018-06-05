const { Client } = require('klasa');

new Client({
  ownerID: process.env.OWNER_ID,
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