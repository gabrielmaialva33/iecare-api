import { inject, injectable } from 'tsyringe'

import { IRole } from 'App/Modules/User/Interfaces/IRole'
import Role from 'App/Modules/User/Models/Role'

import NotFoundException from 'App/Shared/Exceptions/NotFoundException'

@injectable()
export class ShowRoleService {
  constructor(
    @inject('RolesRepository')
    private rolesRepository: IRole.Repository
  ) {}

  public async execute(roleId: string): Promise<Role> {
    const role = await this.rolesRepository.show(roleId)
    if (!role) throw new NotFoundException('Role not found')

    return role
  }
}
