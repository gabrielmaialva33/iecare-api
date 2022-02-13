import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'
import { container } from 'tsyringe'

import {
  CreateProfileService,
  ShowProfilesService,
  UpdateProfileService,
} from 'App/Modules/User/Services/Profile'
import { ProfileValidators } from 'App/Modules/User/Validators/ProfileValidators'

import { ListProfileService } from 'App/Modules/User/Services/Profile/ListProfileService'

export default class ProfilesController {
  public async index({ request, response }: HttpContextContract): Promise<void> {
    const page = request.input('page', 1)
    const perPage = request.input('perPage', 10)
    const search = request.input('search', '')

    const listProfile = container.resolve(ListProfileService)
    const profiles = await listProfile.execute({ page, perPage, search })

    return response.json(profiles)
  }

  public async show({ request, params, response }: HttpContextContract): Promise<void> {
    const { id: profileId } = params

    await request.validate(ProfileValidators.UUID(profileId))

    const showService = container.resolve(ShowProfilesService)
    const profile = await showService.execute(profileId)

    return response.json(profile)
  }

  public async create({ request, response }: HttpContextContract): Promise<void> {
    const profileDto = await request.validate(ProfileValidators.Create)

    const storeService = container.resolve(CreateProfileService)
    const profile = await storeService.execute(profileDto)

    return response.json(profile)
  }

  public async update({ request, params, response }: HttpContextContract): Promise<void> {
    const { id: profileId } = params

    await request.validate(ProfileValidators.UUID(profileId))
    const profileDto = await request.validate(ProfileValidators.Update)

    const updateService = container.resolve(UpdateProfileService)
    const profile = await updateService.execute(profileId, profileDto)

    return response.json(profile)
  }
}
