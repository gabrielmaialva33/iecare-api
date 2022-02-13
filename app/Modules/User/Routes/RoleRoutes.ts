import Route from '@ioc:Adonis/Core/Route'
import RolesController from 'App/Modules/User/Controllers/Http/Admin/RolesController'

Route.group(() => {
  Route.get('/roles', new RolesController().index).as('role.index')
  Route.get('/roles/:id', new RolesController().show).as('role.show')
})
  .prefix('admin')
  .middleware(['auth', 'acl:root,admin'])
