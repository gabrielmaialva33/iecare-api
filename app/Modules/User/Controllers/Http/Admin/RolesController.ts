import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { container } from 'tsyringe'

import { ListRoleService, ShowRoleService } from 'App/Modules/User/Services/Role'
import { RoleValidators } from 'App/Modules/User/Validators/RoleValidators'

export default class RolesController {
  public async index({ request, response }: HttpContextContract) {
    const page = request.input('page', 1)
    const perPage = request.input('perPage', 10)

    const listRoles = container.resolve(ListRoleService)
    const roles = await listRoles.execute(Number(page), Number(perPage))

    return response.json(roles)
  }

  public async show({ request, params, response }: HttpContextContract) {
    const { id: roleId } = params

    await request.validate(RoleValidators.UUID(roleId))

    const showRole = container.resolve(ShowRoleService)
    const role = await showRole.execute(roleId)

    return response.json(role)
  }
}
