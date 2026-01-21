const API_HOST = import.meta.env.VITE_API_HOST || ''

export function buildApiUrl(path, params) {
  const base = API_HOST.endsWith('/') ? API_HOST.slice(0, -1) : API_HOST
  const cleanPath = path.startsWith('/') ? path : `/${path}`
  const url = new URL(`${base}${cleanPath}`, window.location.origin)

  if (params) {
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        url.searchParams.set(key, String(value))
      }
    })
  }

  return url.toString()
}

export async function apiRequest({ method, path }, { params, signal } = {}) {
  const response = await fetch(buildApiUrl(path, params), { method, signal })

  if (!response.ok) {
    throw new Error(`API error: ${response.status}`)
  }

  return response.json()
}
