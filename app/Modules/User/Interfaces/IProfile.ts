import Profile from 'App/Modules/User/Models/Profile'
import { ModelPaginatorContract } from '@ioc:Adonis/Lucid/Orm'

export namespace IProfile {
  export interface Repository {
    list(params: DTO.List): Promise<ModelPaginatorContract<Profile>>

    show(profileId: string): Promise<Profile | null>

    create(data: DTO.Create): Promise<Profile>

    update(profile: Profile): Promise<Profile>
  }

  export interface Helpers {}

  export namespace DTO {
    export interface List {
      page: number
      perPage: number
      search: string
    }

    export interface Create {
      avatar_url?: string
      cellphone: string
      cpf: string
      rg?: string
      company: string
      role: string
      user_id: string
    }

    export interface Update {
      avatar_url?: string
      cellphone?: string
      cpf?: string
      rg?: string
      company?: string
      role?: string
    }
  }
}
