const { Event } = require('klasa');

module.exports = class extends Event {
  async run() {
    this.client.user.setActivity('the Impact Discord', { type: 3 });
  }
};