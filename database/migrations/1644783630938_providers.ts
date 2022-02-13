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

      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
    })
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
