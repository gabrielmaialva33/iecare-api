import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Addresses extends BaseSchema {
  protected tableName = 'addresses'

  public async up() {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

      table.string('cep', 9).nullable()
      table.string('state', 20).nullable()
      table.string('city', 100).nullable()
      table.string('district', 100).nullable()
      table.string('address', 100).nullable()
      table.string('number', 100).nullable()
      table.string('complement', 100).nullable()

      table
        .uuid('provider_id')
        .references('id')
        .inTable('providers')
        .notNullable()
        .onDelete('CASCADE')
        .onUpdate('CASCADE')
        .index('index_addresses_provider_id')

      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
    })
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
