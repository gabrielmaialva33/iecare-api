import { ModelAttributes, ModelPaginatorContract } from '@ioc:Adonis/Lucid/Orm'
import Role from 'App/Modules/User/Models/Role'

export type RoleType = Partial<ModelAttributes<Role>>

export namespace IRole {
  export interface Repository extends Helpers {
    list(page: number, perPage: number): Promise<ModelPaginatorContract<Role>>

    show(roleId: string): Promise<Role | null>
  }

  export interface Helpers {
    findBy(key: string, value: any): Promise<Role | null>
  }
}
