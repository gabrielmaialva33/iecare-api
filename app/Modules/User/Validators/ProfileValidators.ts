import { rules, schema } from '@ioc:Adonis/Core/Validator'
import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export namespace ProfileValidators {
  export class Create {
    constructor(protected ctx: HttpContextContract) {}

    public schema = schema.create({
      avatar_url: schema.string.optional({ trim: true }, [rules.url({})]),
      cellphone: schema.string({ trim: true, escape: true }, [rules.maxLength(30)]),
      cpf: schema.string({ trim: true, escape: true }, [
        rules.unique({ table: 'profiles', column: 'cpf' }),
        rules.minLength(11),
        rules.maxLength(11),
      ]),
      rg: schema.string.optional({ trim: true }, [rules.maxLength(14)]),
      company: schema.string({ trim: true, escape: true }, [rules.maxLength(100)]),
      role: schema.string({ trim: true, escape: true }, [rules.maxLength(100)]),
      user_id: schema.string({ trim: true, escape: true }, [
        rules.uuid({ version: '4' }),
        rules.exists({ table: 'users', column: 'id' }),
        rules.regex(/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i),
      ]),
    })
    public messages = {
      unique: 'Já está em uso.',
      required: 'Não pode ficar em branco.',
      uuid: 'Parâmetro inválido.',
    }
  }

  export class Update {
    constructor(protected ctx: HttpContextContract) {}

    public schema = schema.create({
      avatar_url: schema.string.optional({ trim: true }, [rules.url({})]),
      cellphone: schema.string.optional({ trim: true, escape: true }, [rules.maxLength(30)]),
      cpf: schema.string.optional({ trim: true, escape: true }, [
        rules.unique({ table: 'profiles', column: 'cpf' }),
        rules.minLength(11),
        rules.maxLength(11),
      ]),
      rg: schema.string.optional({ trim: true }, [rules.maxLength(14)]),
      company: schema.string.optional({ trim: true, escape: true }, [rules.maxLength(100)]),
      role: schema.string.optional({ trim: true, escape: true }, [rules.maxLength(100)]),
    })
    public messages = {
      unique: 'Já está em uso.',
      required: 'Não pode ficar em branco.',
      uuid: 'Parâmetro inválido.',
    }
  }

  export function UUID(id: string) {
    return {
      schema: Schemas.UUID,
      data: { id },
      messages: {
        regex: 'Parâmetro invalido',
        uuid: 'Parâmetro invalido',
        exists: 'Perfil está inacessível ou desativado.',
      },
    }
  }

  export function Integer(id: string | number) {
    return {
      schema: Schemas.Integer,
      data: { id },
      messages: {
        number: 'Parâmetro invalido',
        exists: 'Perfil está inacessível ou desativado.',
      },
    }
  }
}

export namespace Schemas {
  export const UUID = schema.create({
    id: schema.string.optional({ trim: true, escape: true }, [
      rules.uuid({ version: '4' }),
      rules.exists({ table: 'profiles', column: 'id' }),
      rules.regex(/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i),
    ]),
  })

  export const Integer = schema.create({
    id: schema.number([rules.unsigned(), rules.exists({ table: 'profiles', column: 'id' })]),
  })
}
