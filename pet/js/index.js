// Imports and uses the 'main.js' generated through gopherjs
'use strict';
console.log('Starting');

console.log('Loading main.js');
require('./main.js');
var main = global.pet;

console.log('Creating a new Pet');
var pet = main.New('brian');
console.log(pet.Name());

console.log('Changing the pet\'s name');
pet.SetName('steve');
console.log(pet.Name());

console.log('Done!');