import { IUser } from 'App/Modules/User/Interfaces/IUser'
import User from 'App/Modules/User/Models/User'

export default class UsersRepository implements IUser.Repository {
  private orm: typeof User
  constructor() {
    this.orm = User
  }

  /**
   * Repository
   */
  public async create(data: IUser.DTO.Create): Promise<User> {
    return this.orm.create(data)
  }

  public async update(user: User): Promise<User> {
    return user.save()
  }

  /**
   * Helpers
   */
  public async findBy(key: string, value: any): Promise<User | null> {
    return this.orm.findBy(key, value)
  }
}
