import { ModelAttributes } from '@ioc:Adonis/Lucid/Orm'
import User from 'App/Modules/User/Models/User'

export type UserType = Partial<ModelAttributes<User>>

export namespace IUser {
  export interface Repository extends Helpers {
    create(data: IUser.DTO.Create): Promise<User>

    update(user: User): Promise<User>
  }

  export interface Helpers {
    findBy(key: string, value: any): Promise<User | null>
  }

  export namespace DTO {
    export interface List {
      page: number
      perPage: number
      search: string
    }

    export interface Create {
      firstname: string
      lastname: string
      email: string
      username: string
      password: string
    }

    export interface Update {
      firstname?: string
      lastname?: string
      email?: string
      username?: string
      password?: string
    }
  }
}
