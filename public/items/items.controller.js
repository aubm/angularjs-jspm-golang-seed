import template from './items.html!text';

export default class ItemsCtrl {

    constructor ($http) {
        "ngInject";

        $http.get('/api/items')
            .then((res) => {
                this.items = res.data;
            });
    }

    static routeConfig($routeProvider) {
        "ngInject";
        $routeProvider.
            when('/items', {
                template: template,
                controller: 'ItemsCtrl',
                controllerAs: 'ic'
            });
    }
}
