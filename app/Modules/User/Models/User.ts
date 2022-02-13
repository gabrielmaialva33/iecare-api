import { DateTime } from 'luxon'
import {
  column,
  beforeSave,
  BaseModel,
  manyToMany,
  ManyToMany,
  beforeCreate,
  afterCreate,
  beforeFind,
  beforeFetch,
  ModelQueryBuilderContract,
} from '@ioc:Adonis/Lucid/Orm'

import Hash from '@ioc:Adonis/Core/Hash'
import Role from 'App/Modules/User/Models/Role'

export default class User extends BaseModel {
  public static table: string = 'users'

  /**
   * ------------------------------------------------------
   * Columns
   * ------------------------------------------------------
   * - column typing struct
   */
  @column({ isPrimary: true })
  public id: string

  @column()
  public firstname: string

  @column()
  public lastname: string

  @column()
  public email: string

  @column()
  public username: string

  @column({ serializeAs: null })
  public password: string

  @column()
  public rememberMeToken?: string

  @column({ serializeAs: null })
  public roleId: string

  @column()
  public isOnline: boolean

  @column({ serializeAs: null })
  public isBlocked: boolean

  @column({ serializeAs: null })
  public isDeleted: boolean

  @column.dateTime({ autoCreate: true, serializeAs: null })
  public createdAt: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true, serializeAs: null })
  public updatedAt: DateTime

  @column.dateTime({ serializeAs: null })
  public deletedAt: DateTime

  /**
   * ------------------------------------------------------
   * Relationships
   * ------------------------------------------------------
   * - define User model relationships
   */
  @manyToMany(() => Role, {
    localKey: 'id',
    pivotForeignKey: 'user_id',
    relatedKey: 'id',
    pivotRelatedForeignKey: 'role_id',
    pivotTable: 'roles_users',
  })
  public roles: ManyToMany<typeof Role>
  /**
   * ------------------------------------------------------
   * Hooks
   * ------------------------------------------------------
   */
  @beforeSave()
  public static async hashPassword(user: User) {
    if (user.$dirty.password) {
      user.password = await Hash.make(user.password)
    }
  }

  @beforeCreate()
  public static async attachDefaultRole(user: User): Promise<void> {
    if (!user.roleId) {
      const userRole = await Role.findBy('name', 'user')
      if (userRole) user.roleId = userRole.id
    }
  }

  @afterCreate()
  public static async attachRoleUser(user: User): Promise<void> {
    if (user.roleId) await user.related('roles').attach([user.roleId])
  }

  @beforeCreate()
  public static async attachUserName(user: User): Promise<void> {
    if (!user.username) {
      user.username = user.email.split('@')[0]
      for (let i = 0; ; i++) {
        const isExists = await User.query().where('username', user.username).first()
        if (isExists) user.username = `${user.username}${i}`
        else break
      }
    }
  }

  @beforeFind()
  @beforeFetch()
  public static async ignoreDeleted(query: ModelQueryBuilderContract<typeof User>): Promise<void> {
    query.whereNot('is_deleted', true)
  }

  /**
   * ------------------------------------------------------
   * Query Scopes
   * ------------------------------------------------------
   */

  /**
   * ------------------------------------------------------
   * Custom Methods
   * ------------------------------------------------------
   */
}
