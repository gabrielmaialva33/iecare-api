import { rules, schema } from '@ioc:Adonis/Core/Validator'

export namespace RoleValidators {
  export function UUID(id: string) {
    return {
      schema: Schemas.UUID,
      data: { id },
      messages: {
        regex: 'Parâmetro invalido',
        uuid: 'Parâmetro invalido',
        exists: 'Regra está inacessível ou desativada.',
      },
    }
  }

  export function Integer(id: string | number) {
    return {
      schema: Schemas.Integer,
      data: { id },
      messages: {
        number: 'Parâmetro invalido',
        exists: 'Regra está inacessível ou desativada.',
      },
    }
  }
}

export namespace Schemas {
  export const UUID = schema.create({
    id: schema.string.optional({ trim: true, escape: true }, [
      rules.uuid({ version: '4' }),
      rules.exists({ table: 'roles', column: 'id', whereNot: { name: 'root' } }),
      rules.regex(/^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i),
    ]),
  })

  export const Integer = schema.create({
    id: schema.number([
      rules.unsigned(),
      rules.exists({ table: 'roles', column: 'id', whereNot: { name: 'root' } }),
    ]),
  })
}
