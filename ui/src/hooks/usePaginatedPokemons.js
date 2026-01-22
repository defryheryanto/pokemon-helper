import { useEffect, useState } from 'react'
import { fetchPokemons } from '../services/pokemonService'

function usePaginatedPokemons({ pageSize = 50, search = '', elementType = '' } = {}) {
  const [pokemons, setPokemons] = useState([])
  const [page, setPage] = useState(1)
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState(null)
  const [hasMore, setHasMore] = useState(true)

  useEffect(() => {
    setPage(1)
    setPokemons([])
    setHasMore(true)
  }, [search, elementType, pageSize])

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
          search,
          elementType,
          signal: controller.signal,
        })

        if (!isMounted) {
          return
        }

        setPokemons(data)
        setHasMore(data.length === pageSize)
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
  }, [page, pageSize, search, elementType])

  return {
    pokemons,
    page,
    setPage,
    isLoading,
    error,
    hasMore,
  }
}

export default usePaginatedPokemons
