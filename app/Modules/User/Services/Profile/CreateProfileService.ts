import { inject, injectable } from 'tsyringe'

import { IProfile } from 'App/Modules/User/Interfaces/IProfile'
import Profile from 'App/Modules/User/Models/Profile'

@injectable()
export class CreateProfileService {
  constructor(
    @inject('ProfilesRepository')
    private profilesRepository: IProfile.Repository
  ) {}

  public async execute(data: IProfile.DTO.Create): Promise<Profile> {
    return this.profilesRepository.create(data)
  }
}
