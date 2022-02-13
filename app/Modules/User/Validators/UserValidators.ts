import { rules, schema } from '@ioc:Adonis/Core/Validator'
import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

export namespace UserValidators {
  export class Create {
    constructor(protected ctx: HttpContextContract) {}

    public schema = schema.create({
      firstname: schema.string({ trim: true, escape: true }, [rules.maxLength(80)]),
      lastname: schema.string({ trim: true, escape: true }, [rules.maxLength(80)]),
      username: schema.string({ trim: true }, [
        rules.unique({ table: 'users', column: 'username' }),
        rules.minLength(4),
        rules.maxLength(80),
      ]),
      email: schema.string({ trim: true }, [
        rules.email(),
        rules.unique({ table: 'users', column: 'email' }),
      ]),
      password: schema.string({ trim: true }, [rules.confirmed(), rules.minLength(6)]),
    })
    public messages = {
      email: 'E-mail não é valido.',
      unique: 'Já está em uso.',
      minLength: 'Requer mais que 4 caracteres.',
      maxLength: 'Não pode ter mais que 80 caracteres.',
      required: 'Não pode ficar em branco.',
      uuid: 'Parâmetro inválido.',
    }
  }

  export class Update {
    constructor(protected ctx: HttpContextContract) {}

    public schema = schema.create({
      firstname: schema.string.optional({ trim: true, escape: true }, [rules.maxLength(80)]),
      lastname: schema.string.optional({ trim: true, escape: true }, [rules.maxLength(80)]),
      username: schema.string.optional({ trim: true }, [
        rules.unique({ table: 'users', column: 'username' }),
        rules.minLength(4),
        rules.maxLength(80),
      ]),
      email: schema.string.optional({ trim: true }, [
        rules.email(),
        rules.unique({ table: 'users', column: 'email' }),
      ]),
      password: schema.string.optional({ escape: true }, [rules.confirmed(), rules.minLength(6)]),
    })
    public messages = {
      email: 'E-mail não é valido.',
      unique: 'Já está em uso.',
      minLength: 'Requer mais que 4 caracteres.',
      maxLength: 'Não pode ter mais que 80 caracteres.',
      required: 'Não pode ficar em branco.',
      uuid: 'Parâmetro inválido.',
    }
  }

  export class Login {
    constructor(protected ctx: HttpContextContract) {}

    public schema = schema.create({
      uid: schema.string({ trim: true }, []),
      password: schema.string({ trim: true }),
    })

    public messages = {
      required: 'Não pode ficar em branco.',
    }
  }

  export function UUID(id: string) {
    return {
      schema: Schemas.UUID,
      data: { id },
      messages: {
        regex: 'Parâmetro invalido',
        uuid: 'Parâmetro invalido',
        exists: 'Usuário está inacessível ou desativado.',
      },
    }
  }

  export function Integer(id: string | number) {
    return {
      schema: Schemas.Integer,
      data: { id },
      messages: {
        number: 'Parâmetro invalido',
        exists: 'Usuário está inacessível ou desativado.',
      },
    }
  }
}

export namespace Schemas {
  export const UUID = schema.create({
    id: schema.string.optional({ trim: true, escape: true }, [
      rules.uuid({ version: '4' }),
      rules.exists({ table: 'users', column: 'id', whereNot: { is_deleted: true } }),
      rules.regex(/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i),
    ]),
  })

  export const Integer = schema.create({
    id: schema.number([
      rules.unsigned(),
      rules.exists({ table: 'users', column: 'id', whereNot: { is_deleted: true } }),
    ]),
  })
}
