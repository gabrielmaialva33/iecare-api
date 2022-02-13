import { container, delay } from 'tsyringe'

import { IRole } from 'App/Modules/User/Interfaces/IRole'
import RolesRepository from 'App/Modules/User/Repositories/RolesRepository'

import { IUser } from 'App/Modules/User/Interfaces/IUser'
import UsersRepository from 'App/Modules/User/Repositories/UsersRepository'

import { IProfile } from 'App/Modules/User/Interfaces/IProfile'
import ProfilesRepository from 'App/Modules/User/Repositories/ProfilesRepository'

container.registerSingleton<IRole.Repository>(
  'RolesRepository',
  delay(() => RolesRepository)
)

container.registerSingleton<IUser.Repository>(
  'UsersRepository',
  delay(() => UsersRepository)
)

container.registerSingleton<IProfile.Repository>(
  'ProfilesRepository',
  delay(() => ProfilesRepository)
)
