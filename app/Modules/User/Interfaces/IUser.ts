import { ModelAttributes } from '@ioc:Adonis/Lucid/Orm'
import User from 'App/Modules/User/Models/User'

export type UserType = Partial<ModelAttributes<User>>

export namespace IUser {
  export interface Repository {}

  export interface Helpers {}

  export namespace DTO {
    export interface List {}

    export interface Create {}

    export interface Update {}
  }
}
