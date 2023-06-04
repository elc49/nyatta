import { type CookieValueTypes } from 'cookies-next'
import { setContext } from '@apollo/client/link/context'

const authLink = (jwt?: CookieValueTypes) => setContext((_, previousContext) => {
  const { headers } = previousContext
  return {
    ...previousContext,
    headers: {
      ...headers,
      'keep-alive': 'true',
      ...(jwt && { Authorization: `Bearer ${jwt}` })
    }
  }
})

export default authLink
