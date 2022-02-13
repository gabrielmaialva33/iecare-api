import Profile from 'App/Modules/User/Models/Profile'

export namespace IProfile {
  export interface Repository {
    show(profileId: string): Promise<Profile | null>

    create(data: DTO.Store): Promise<Profile>

    update(profile: Profile): Promise<Profile>
  }

  export interface Helpers {}

  export namespace DTO {
    export interface Store {
      avatarUrl?: string
      cellphone: string
      cpf: string
      rg?: string
      company: string
      role: string
      userId: string
    }

    export interface Update {
      avatarUrl?: string
      cellphone?: string
      cpf?: string
      rg?: string
      company?: string
      role?: string
    }
  }
}
