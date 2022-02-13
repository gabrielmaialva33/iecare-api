import { injectable, inject } from 'tsyringe'
import { ModelPaginatorContract } from '@ioc:Adonis/Lucid/Orm'

import { IRole } from 'App/Modules/User/Interfaces/IRole'
import Role from 'App/Modules/User/Models/Role'

@injectable()
export class ListRoleService {
  constructor(
    @inject('RolesRepository')
    private rolesRepository: IRole.Repository
  ) {}

  public async execute(page: number, perPage: number): Promise<ModelPaginatorContract<Role>> {
    return this.rolesRepository.list(page, perPage)
  }
}
