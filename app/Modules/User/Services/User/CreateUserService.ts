import { inject, injectable } from 'tsyringe'

import User from 'App/Modules/User/Models/User'
import { IUser } from 'App/Modules/User/Interfaces/IUser'

@injectable()
export class CreateUserService {
  constructor(
    @inject('UsersRepository')
    private usersRepository: IUser.Repository
  ) {}

  public async execute(data: IUser.DTO.Create): Promise<User> {
    return this.usersRepository.create(data)
  }
}
