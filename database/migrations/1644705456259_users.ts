import BaseSchema from '@ioc:Adonis/Lucid/Schema'
import Logger from '@ioc:Adonis/Core/Logger'

export default class UsersSchema extends BaseSchema {
  protected tableName = 'users'

  public async up() {
    if (!(await this.schema.hasTable(this.tableName)))
      this.schema.createTable(this.tableName, (table) => {
        table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

        table.string('firstname', 100).notNullable().index('index_user_firstname')
        table.string('lastname', 100).notNullable().index('index_user_lastname')
        table.string('email', 255).notNullable().index('index_user_email')
        table.string('username', 50).notNullable().index('index_user_username')
        table.string('password', 180).notNullable()

        table.string('remember_me_token').nullable()

        table
          .uuid('role_id')
          .references('id')
          .inTable('roles')
          .notNullable()
          .onDelete('CASCADE')
          .onUpdate('CASCADE')
          .index('index_user_role_id')

        table.boolean('is_online').notNullable().defaultTo(false)
        table.boolean('is_blocked').notNullable().defaultTo(false)
        table.boolean('is_deleted').notNullable().defaultTo(false)

        table.timestamp('created_at', { useTz: true }).notNullable()
        table.timestamp('updated_at', { useTz: true }).notNullable()
        table.timestamp('deleted_at', { useTz: true }).nullable()
      })
    else Logger.info('User migration already running')
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
