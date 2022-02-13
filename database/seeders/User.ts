import BaseSeeder from '@ioc:Adonis/Lucid/Seeder'
import Logger from '@ioc:Adonis/Core/Logger'

import User from 'App/Modules/User/Models/User'
import Role from 'App/Modules/User/Models/Role'

export default class UserSeeder extends BaseSeeder {
  public async run() {
    try {
      const users = await User.query().whereIn('username', ['root', 'admin'])
      if (users.length) return Logger.info('User already sown.')

      const root = await Role.findBy('name', 'root')
      const admin = await Role.findBy('name', 'admin')
      const user = await Role.findBy('name', 'user')
      const guest = await Role.findBy('name', 'guest')

      if (admin && root && user && guest)
        await User.createMany([
          {
            firstname: 'Root',
            lastname: 'User',
            email: 'root@iecare.com.br',
            username: 'root',
            password: 'iecare@551238',
            roleId: root.id,
          },
          {
            firstname: 'Admin',
            lastname: 'User',
            email: 'admin@iecare.com.br',
            username: 'admin',
            password: 'iecare@551238',
            roleId: admin.id,
          },
          {
            firstname: 'Gabriel',
            lastname: 'Maia',
            email: 'gabriel.maia@iecare.com.br',
            username: 'gabriel.maia',
            password: 'iecare@551238',
            roleId: user.id,
          },
          {
            firstname: 'Guest',
            lastname: 'User',
            email: 'guest@iecare.com.br',
            username: 'guest',
            password: 'iecare@551238',
            roleId: guest.id,
          },
        ])
    } catch (err) {
      Logger.error(err.message)
    }
  }
}
