import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Providers extends BaseSchema {
  protected tableName = 'providers'

  public async up() {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

      table.string('name', 100).notNullable()
      table.string('cellphone', 30).notNullable()
      table.string('cnpj', 14).nullable()

      table.boolean('is_mei').notNullable().defaultTo(false)
      table.boolean('is_me').notNullable().defaultTo(false)

      table
        .uuid('user_id')
        .references('id')
        .inTable('users')
        .notNullable()
        .onDelete('CASCADE')
        .onUpdate('CASCADE')
        .index('index_provider_user_id')

      table.boolean('is_deleted').notNullable().defaultTo(false)

      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
      table.timestamp('deleted_at', { useTz: true }).nullable()
    })
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
