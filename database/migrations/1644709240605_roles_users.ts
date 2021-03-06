import BaseSchema from '@ioc:Adonis/Lucid/Schema'
import Logger from '@ioc:Adonis/Core/Logger'

export default class RolesUsers extends BaseSchema {
  protected tableName = 'roles_users'

  public async up() {
    if (!(await this.schema.hasTable(this.tableName)))
      this.schema.createTable(this.tableName, (table) => {
        table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

        table
          .uuid('role_id')
          .references('id')
          .inTable('roles')
          .notNullable()
          .onDelete('CASCADE')
          .onUpdate('CASCADE')

        table
          .uuid('user_id')
          .references('id')
          .inTable('users')
          .notNullable()
          .onDelete('CASCADE')
          .onUpdate('CASCADE')

        table.timestamps(true, true)
      })
    else Logger.info('RolesUsers migration already running')
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
