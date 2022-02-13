import { IRole } from 'App/Modules/User/Interfaces/IRole'
import Role from 'App/Modules/User/Models/Role'
import { ModelPaginatorContract } from '@ioc:Adonis/Lucid/Orm'

export default class RolesRepository implements IRole.Repository {
  private orm: typeof Role
  constructor() {
    this.orm = Role
  }

  /**
   * Repository
   */
  public async list(page: number, perPage: number): Promise<ModelPaginatorContract<Role>> {
    return this.orm.query().whereNot({ name: 'root' }).paginate(page, perPage)
  }

  public async show(roleId: string): Promise<Role | null> {
    return this.orm.query().where({ id: roleId }).first()
  }
}
