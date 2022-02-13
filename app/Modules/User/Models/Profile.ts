import { BaseModel, BelongsTo, belongsTo, column, scope } from '@ioc:Adonis/Lucid/Orm'
import { DateTime } from 'luxon'

import User from 'App/Modules/User/Models/User'

export default class Profile extends BaseModel {
  public static table: string = 'profiles'

  /**
   * ------------------------------------------------------
   * Columns
   * ------------------------------------------------------
   * - column typing struct
   */
  @column({ isPrimary: true })
  public id: string

  @column()
  public avatarUrl: string

  @column()
  public cellphone: string

  @column()
  public cpf: string

  @column()
  public rg: string

  @column()
  public company: string

  @column()
  public role: string

  @column({ serializeAs: null })
  public user_id: string

  @column.dateTime({ autoCreate: true, serializeAs: null })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true, serializeAs: null })
  public updatedAt: DateTime

  /**
   * ------------------------------------------------------
   * Relationships
   * ------------------------------------------------------
   * - define account settings  model relationships
   * */
  @belongsTo(() => User, { localKey: 'id', foreignKey: 'user_id' })
  public user: BelongsTo<typeof User>

  /**
   * ------------------------------------------------------
   * Query Scopes
   * ------------------------------------------------------
   */

  public static searchQueryScope = scope((query, search) => {
    const fields = ['cellphone', 'cpf', 'rg', 'company', 'role']
    let sql = ''

    fields.forEach((field, i) => {
      sql = `${sql} ${i !== 0 ? ' or ' : ' '} ${field} like '%${search}%'`
    })

    return query.whereRaw(`(${sql})`)
  })
}
