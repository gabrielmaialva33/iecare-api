import { HttpContextContract } from '@ioc:Adonis/Core/HttpContext'

import { UserValidators } from 'App/Modules/User/Validators/UserValidators'

import AuthorizationException from 'App/Shared/Exceptions/AuthorizationException'

export default class AuthController {
  public async login({ request, auth, response }: HttpContextContract) {
    const { uid, password } = await request.validate(UserValidators.Login)

    try {
      const token = await auth
        .use('api')
        .attempt(uid, password, { name: 'iecare-token', expiresIn: '1h' })

      return response.json({ user: auth.user, token })
    } catch (error) {
      throw new AuthorizationException(
        'Não foi possível fazer o login, verifique suas credenciais ou tente novamente mais tarte.'
      )
    }
  }
}
