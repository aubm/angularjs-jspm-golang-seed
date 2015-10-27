import angular from 'angular';
import Func1 from './messages/messages';
import Items from './items/items';
import Config from './config/config';

let app = angular.module('app', [
  Func1.name,
  Config.name,
  Items.name
]);

angular.element(document).ready(() =>  angular.bootstrap(document.body, [ app.name ], { strictDi: false }));
