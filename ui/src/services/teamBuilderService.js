import { apiRequest } from './apiClient'
import { API_ENDPOINTS } from './apiEndpoints'

export async function fetchTeamSimulation(payload, { signal } = {}) {
  return apiRequest(API_ENDPOINTS.simulateTeam, {
    body: payload,
    signal,
  })
}
