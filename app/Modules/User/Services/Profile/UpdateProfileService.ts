import { inject, injectable } from 'tsyringe'

import { IProfile } from 'App/Modules/User/Interfaces/IProfile'
import Profile from 'App/Modules/User/Models/Profile'
import NotFoundException from 'App/Shared/Exceptions/NotFoundException'

@injectable()
export class UpdateProfileService {
  constructor(
    @inject('ProfilesRepository')
    private profilesRepository: IProfile.Repository
  ) {}

  public async execute(profileId: string, data: IProfile.DTO.Update): Promise<Profile> {
    const profile = await this.profilesRepository.show(profileId)
    if (!profile) throw new NotFoundException('Profile not found')

    profile.merge(data)
    await this.profilesRepository.update(profile)

    return profile
  }
}
