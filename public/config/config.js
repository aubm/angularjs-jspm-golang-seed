import angular from 'angular';
import Bootstrap from './bootstrap/bootstrap';
import RouteConfig from './route';

export default angular.module('app.config', [
  Bootstrap.name,
  RouteConfig.name
]);