import { inject, injectable } from 'tsyringe'
import { ModelPaginatorContract } from '@ioc:Adonis/Lucid/Orm'

import { IProfile } from 'App/Modules/User/Interfaces/IProfile'
import Profile from 'App/Modules/User/Models/Profile'

@injectable()
export class ListProfileService {
  constructor(
    @inject('ProfilesRepository')
    private profilesRepository: IProfile.Repository
  ) {}

  public async execute({
    page,
    perPage,
    search,
  }: IProfile.DTO.List): Promise<ModelPaginatorContract<Profile>> {
    return this.profilesRepository.list({ page, perPage, search })
  }
}
