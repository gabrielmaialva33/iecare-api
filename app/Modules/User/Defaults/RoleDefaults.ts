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
    slug: 'Provider',
    name: 'provider',
    description: 'a provider user system',
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
