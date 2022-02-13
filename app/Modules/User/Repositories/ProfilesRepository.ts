import { IProfile } from 'App/Modules/User/Interfaces/IProfile'
import Profile from 'App/Modules/User/Models/Profile'

export default class ProfilesRepository implements IProfile.Repository {
  private orm: typeof Profile

  constructor() {
    this.orm = Profile
  }
  /**
   * Repository
   */
  public async show(profileId: string): Promise<Profile | null> {
    return this.orm.query().where({ id: profileId }).first()
  }

  public async create(data: IProfile.DTO.Store): Promise<Profile> {
    return this.orm.create(data)
  }

  public async update(profile: Profile): Promise<Profile> {
    return profile.save()
  }

  /**
   * Helpers
   */
}
