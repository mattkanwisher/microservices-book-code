app = angular.module('user', ['ngRoute','webAdminApi'])

.config(function($routeProvider) {
  $routeProvider
    .when('/', {
      controller:'UserListController as userList',
      templateUrl:'/assets/views/list.html'
    })
    .when('/new', {
      controller:'NewUserController as newUser',
      templateUrl:'/assets/views/new.html'
    })
    .when('/edit/:userId', {
      controller:'EditUserController as editUser',
      templateUrl:'/assets/views/edit.html'
    })
    .when('/view/:userId', {
      controller:'ViewUserController as viewUser',
      templateUrl:'/assets/views/view.html'
    })
    .otherwise({
      redirectTo:'/'
    });
})

.controller('UserListController', function($scope, userApi) {
  var userList = this;
  $scope.users = userApi.users();
  $scope.delete = function(index, id){
    $scope.users.splice(index, 1)
    userApi.remove({userId:id});
  }
})

.controller('NewUserController', function($scope, $location, userApi) {
  $scope.create = function(){
    userApi.create($scope.user, function(){
      $location.path("/");
    },function(resp){
      if (resp.status == 400) {
        $scope.validate = resp.data.message;
      }else{
        $scope.error = resp.data.message;
      }
    });
  }
})

.controller('EditUserController', function($scope, $routeParams, $location, userApi) {
  $scope.info = {};
  $scope.pwd = {};

  $scope.user = userApi.user({userId:$routeParams.userId}, function(){
    $scope.info = {
      id: $scope.user.id,
      name: $scope.user.name,
      email: $scope.user.email
    }

    $scope.pwd = {
      id: $scope.user.id
    }
  },function(){
    $location.path("/")
  });

  $scope.infoUpdate = function(){
    userApi.edit($scope.info, function(){
      $location.path("/")
    },function(resp){
      if (resp.status == 400) {
        console.log(resp.data.message);
        $scope.infoValidate = resp.data.message;
      }else{
        $scope.infoError = resp.data.message;
      }
    })
  }

  $scope.pwdUpdate = function(){
    userApi.pwd($scope.pwd, function(){
      $location.path("/")
    }, function(resp){
      if (resp.status == 400) {
        $scope.pwdValidate = resp.data.message;
      }else{
        $scope.pwdError = resp.data.message;
      }
    })
  }
})

.controller('ViewUserController', function($scope, $routeParams, $location, userApi) {
  $scope.user = userApi.user({userId:$routeParams.userId})
})

