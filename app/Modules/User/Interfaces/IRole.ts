import { ModelAttributes } from '@ioc:Adonis/Lucid/Orm'
import Role from 'App/Modules/User/Models/Role'

export type RoleType = Partial<ModelAttributes<Role>>

export namespace IRole {
  export interface Repository {}

  export interface Helpers {}

  export namespace DTO {
    export interface List {}

    export interface Create {}

    export interface Update {}
  }
}
