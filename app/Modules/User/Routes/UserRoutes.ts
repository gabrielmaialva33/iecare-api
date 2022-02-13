import Route from '@ioc:Adonis/Core/Route'
import UsersController from 'App/Modules/User/Controllers/Http/User/UsersController'
import AuthController from 'App/Modules/User/Controllers/Http/User/AuthController'

Route.post('login', new AuthController().login).as('auth.login')
Route.post('/users', new UsersController().create).as('user.create')

Route.group(() => {
  Route.get('/users/:id', new UsersController().show).as('user.show')
  Route.put('/users/:id', new UsersController().update).as('user.update')
}).middleware(['auth', 'acl:root,admin,user,guest'])
