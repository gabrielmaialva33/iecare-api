import Route from '@ioc:Adonis/Core/Route'
import ProfilesController from 'App/Modules/User/Controllers/Http/User/ProfilesController'

Route.group(() => {
  Route.get('/profiles', new ProfilesController().index).as('profile.index')
  Route.get('/profiles/:id', new ProfilesController().show).as('profile.show')
  Route.post('/profiles', new ProfilesController().create).as('profile.create')
  Route.put('/profiles/:id', new ProfilesController().update).as('profile.update')
}).middleware(['auth', 'acl:root,admin,user,guest'])
