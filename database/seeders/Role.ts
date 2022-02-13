import BaseSeeder from '@ioc:Adonis/Lucid/Seeder'
import Logger from '@ioc:Adonis/Core/Logger'

import Role from 'App/Modules/User/Models/Role'
import { RoleDefaults } from 'App/Modules/User/Defaults/RoleDefaults'

export default class RoleSeeder extends BaseSeeder {
  public async run(): Promise<void> {
    try {
      const roles = await Role.query().whereIn('name', ['root', 'admin', 'user'])
      if (roles.length) return Logger.info('Role already sown.')
      await Role.createMany(RoleDefaults)
    } catch (err) {
      Logger.error(err.message)
    }
  }
}
