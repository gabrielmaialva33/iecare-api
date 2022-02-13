import { RoleType } from 'App/Modules/User/Interfaces/IRole'

export const RoleDefaults: RoleType[] = [
  {
    slug: 'Root',
    name: 'root',
    description: 'root system',
  },
  {
    slug: 'Admin',
    name: 'admin',
    description: 'admin system',
  },
  {
    slug: 'User',
    name: 'user',
    description: 'a common user system',
  },
  {
    slug: 'Guest',
    name: 'guest',
    description: 'a guest user system',
  },
]
