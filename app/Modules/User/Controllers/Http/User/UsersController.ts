import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { container } from 'tsyringe'

import {
  ShowUserService,
  CreateUserService,
  UpdateUserService,
} from 'App/Modules/User/Services/User'
import { UserValidators } from 'App/Modules/User/Validators/UserValidators'

export default class UsersController {
  public async index({}: HttpContextContract) {}

  public async show({ request, params, response }: HttpContextContract) {
    const { id: userId } = params

    await request.validate(UserValidators.UUID(userId))

    const showService = container.resolve(ShowUserService)
    const user = await showService.execute(userId)

    return response.json(user)
  }

  public async create({ request, response }: HttpContextContract) {
    const userDto = await request.validate(UserValidators.Create)

    const storeService = container.resolve(CreateUserService)
    const user = await storeService.execute(userDto)

    return response.json(user)
  }

  public async update({ request, params, response }: HttpContextContract) {
    const { id: userId } = params

    await request.validate(UserValidators.UUID(userId))

    const userDto = await request.validate(UserValidators.Update)

    const updateService = container.resolve(UpdateUserService)
    const user = await updateService.execute(userId, userDto)

    return response.json(user)
  }

  public async destroy({}: HttpContextContract) {}
}
