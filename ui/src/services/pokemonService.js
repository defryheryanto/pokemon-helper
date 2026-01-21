import { apiRequest } from './apiClient'
import { API_ENDPOINTS } from './apiEndpoints'

export async function fetchPokemons({ page = 1, pageSize = 50, signal } = {}) {
  const payload = await apiRequest(API_ENDPOINTS.pokemons, {
    params: { page, pageSize },
    signal,
  })

  return payload?.pokemons ?? []
}
