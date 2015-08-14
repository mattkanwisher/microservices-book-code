app = angular.module('webAdminApi', ['ngResource'])

.value('apiURL', 'http://127.0.0.1:3000')

.service('userApi', function($resource, apiURL) {
  return $resource(apiURL + '/user/:userId/:action', {userId:'@id'}, {
    users: {method:'GET', isArray:true},
    user: {method:'GET'},
    remove: {method:'DELETE'},
    create: {method:'PUT'},
    edit: {method:'POST'},
    pwd: {method:'POST', params:{action:'password'}}
  });
})