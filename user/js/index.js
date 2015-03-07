// Imports and uses the 'main.js' generated through gopherjs
'use strict';
console.log('Starting');

console.log('Loading main.js');
require('./main.js'); // Attaches methods to a 'user' property on the 'global' object
var main = global.user;

console.log('Creating a DB adapter');
var data = [
  ['Jane', 1],
  ['John', 2],
  ['Sarah', 3],
  ['Steve', 4],
  ['Jess', 5]
];
var db = {
  Query: function(query, params) {
    // Cheating a bit here...
    switch (true) {
    case query.indexOf('UPDATE') === 0:
      data[params[1] - 1][0] = params[0];
      break;
    case query.indexOf('INSERT') === 0:
      data.push([params[0], data.length + 1]);
      break;
    case query === 'SELECT @@IDENTITY':
      return [[data.length]];
      break;
    case query.indexOf('WHERE') !== -1:
      return [data[params[0] - 1]];
      break;
    case query.indexOf('SELECT') === 0:
      return data;
      break;
    }
    return [];
  }
};

console.log('Registering the DB adapter');
main.registerDB(db);

console.log('Getting all users');
console.log(main.all());

console.log('Adding a new user');
main.new('Richard');
console.log(main.get(6));

console.log('Update the name for a user');
var user = main.get(2);
user.Name = 'Jason';
main.save(user);
console.log(main.all());

console.log('Done!');