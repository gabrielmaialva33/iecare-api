import BaseSchema from '@ioc:Adonis/Lucid/Schema'
import Logger from '@ioc:Adonis/Core/Logger'

export default class Roles extends BaseSchema {
  protected tableName = 'roles'

  public async up() {
    if (!(await this.schema.hasTable(this.tableName)))
      this.schema.createTable(this.tableName, (table) => {
        table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

        table.string('slug', 100).notNullable().index('index_role_slug')
        table.string('name', 100).notNullable().unique().index('index_role_name')
        table.string('description', 100).notNullable()
        table.jsonb('permissions').defaultTo('{}')

        table.timestamp('created_at', { useTz: true })
        table.timestamp('updated_at', { useTz: true })
      })
    else Logger.info('Roles migration already running')
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
