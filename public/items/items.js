import angular from 'angular';
import ItemsCtrl from './items.controller';
import AppRouteConfig from '../config/route';

export default angular
  .module('app.items', [
    AppRouteConfig.name
  ])
    .controller('ItemsCtrl', ItemsCtrl)
    .config(ItemsCtrl.routeConfig);
