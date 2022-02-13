import { inject, injectable } from 'tsyringe'

import { IUser } from 'App/Modules/User/Interfaces/IUser'
import User from 'App/Modules/User/Models/User'

import NotFoundException from 'App/Shared/Exceptions/NotFoundException'

@injectable()
export class ShowUserService {
  constructor(
    @inject('UsersRepository')
    private usersRepository: IUser.Repository
  ) {}

  public async execute(userId: string): Promise<User> {
    const user = await this.usersRepository.findBy('id', userId)
    if (!user) throw new NotFoundException('User not found')

    return user
  }
}
