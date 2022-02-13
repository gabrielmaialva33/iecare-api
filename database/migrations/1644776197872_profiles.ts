import BaseSchema from '@ioc:Adonis/Lucid/Schema'

export default class Profiles extends BaseSchema {
  protected tableName = 'profiles'

  public async up() {
    this.schema.createTable(this.tableName, (table) => {
      table.uuid('id').primary().defaultTo(this.db.rawQuery('uuid_generate_v4()').knexQuery)

      table.string('avatar_url').nullable()
      table.string('cellphone', 30).nullable()
      table.string('cpf', 11).nullable()
      table.string('rg', 14).nullable()

      table.string('company', 100).nullable()
      table.string('role', 50).nullable()

      table
        .uuid('user_id')
        .references('id')
        .inTable('users')
        .unique()
        .notNullable()
        .onDelete('CASCADE')
        .onUpdate('CASCADE')
        .index('index_profile_user_id')

      table.timestamp('created_at', { useTz: true })
      table.timestamp('updated_at', { useTz: true })
    })
  }

  public async down() {
    this.schema.dropTable(this.tableName)
  }
}
