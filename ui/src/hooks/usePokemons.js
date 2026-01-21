import { useEffect, useState } from 'react'
import { fetchPokemons } from '../services/pokemonService'

function usePokemons({ page = 1, pageSize = 50 } = {}) {
  const [pokemons, setPokemons] = useState([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    const controller = new AbortController()
    let isMounted = true

    async function loadPokemons() {
      setIsLoading(true)
      setError(null)
      try {
        const data = await fetchPokemons({
          page,
          pageSize,
          signal: controller.signal,
        })
        if (isMounted) {
          setPokemons(data)
        }
      } catch (err) {
        if (err.name === 'AbortError') {
          return
        }
        if (isMounted) {
          setError(err)
        }
      } finally {
        if (isMounted) {
          setIsLoading(false)
        }
      }
    }

    loadPokemons()

    return () => {
      isMounted = false
      controller.abort()
    }
  }, [page, pageSize])

  return { pokemons, isLoading, error }
}

export default usePokemons
