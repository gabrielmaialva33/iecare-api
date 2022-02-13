import { container, delay } from 'tsyringe'

import { IRole } from 'App/Modules/User/Interfaces/IRole'
import RolesRepository from 'App/Modules/User/Repositories/RolesRepository'

import { IUser } from 'App/Modules/User/Interfaces/IUser'
import UsersRepository from 'App/Modules/User/Repositories/UsersRepository'

container.registerSingleton<IRole.Repository>(
  'RolesRepository',
  delay(() => RolesRepository)
)

container.registerSingleton<IUser.Repository>(
  'UsersRepository',
  delay(() => UsersRepository)
)
