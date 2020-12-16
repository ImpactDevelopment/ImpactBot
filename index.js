const Discord = require('discord.js');
const client = new Discord.Client();

					client.on(
							'message',
		msg => {
if   ( msg.author.bot    )  {
	return;

					}   else {
msg.channel.send('iam impactt bot');
					}


					}
			)
			;

																							client
	.login(
									"token"
						);