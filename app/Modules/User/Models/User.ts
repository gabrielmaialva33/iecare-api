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
  computed,
  hasOne,
  HasOne,
  hasMany,
  HasMany,
} from '@ioc:Adonis/Lucid/Orm'

import Hash from '@ioc:Adonis/Core/Hash'
import Role from 'App/Modules/User/Models/Role'
import Profile from 'App/Modules/User/Models/Profile'
import Provider from 'App/Modules/Provider/Models/Provider'

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
  public remember_me_token?: string

  @column({ serializeAs: null })
  public role_id: string

  @column()
  public is_online: boolean

  @column({ serializeAs: null })
  public is_blocked: boolean

  @column({ serializeAs: null })
  public is_deleted: boolean

  @column.dateTime({ autoCreate: true, serializeAs: null })
  public created_at: DateTime

  @column.dateTime({ autoCreate: true, autoUpdate: true, serializeAs: null })
  public updated_at: DateTime

  @column.dateTime({ serializeAs: null })
  public deleted_at: DateTime

  /**
   * ------------------------------------------------------
   * Computed
   * ------------------------------------------------------
   */
  @computed()
  public get fullname() {
    return `${this.firstname} ${this.lastname}`
  }

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

  @hasOne(() => Profile, { localKey: 'id', foreignKey: 'user_id' })
  public profile: HasOne<typeof Profile>

  @hasMany(() => Provider, { localKey: 'id', foreignKey: 'user_id' })
  public providers: HasMany<typeof Provider>

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
    if (!user.role_id) {
      const userRole = await Role.findBy('name', 'user')
      if (userRole) user.role_id = userRole.id
    }
  }

  @afterCreate()
  public static async attachRoleUser(user: User): Promise<void> {
    if (user.role_id) await user.related('roles').attach([user.role_id])
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
