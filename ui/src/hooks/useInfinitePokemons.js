import { useCallback, useEffect, useRef, useState } from 'react'
import { fetchPokemons } from '../services/pokemonService'

function useInfinitePokemons({ pageSize = 50 } = {}) {
  const [pokemons, setPokemons] = useState([])
  const [page, setPage] = useState(1)
  const [isLoading, setIsLoading] = useState(true)
  const [isLoadingMore, setIsLoadingMore] = useState(false)
  const [error, setError] = useState(null)
  const [hasMore, setHasMore] = useState(true)
  const inFlightRef = useRef(false)

  const loadMore = useCallback(() => {
    if (inFlightRef.current || !hasMore) {
      return
    }
    setPage((prev) => prev + 1)
  }, [hasMore])

  useEffect(() => {
    const controller = new AbortController()
    let isMounted = true

    async function loadPokemons() {
      const isFirstPage = page === 1
      inFlightRef.current = true
      setError(null)
      if (isFirstPage) {
        setIsLoading(true)
      } else {
        setIsLoadingMore(true)
      }

      try {
        const data = await fetchPokemons({
          page,
          pageSize,
          signal: controller.signal,
        })

        if (!isMounted) {
          return
        }

        setPokemons((prev) => (isFirstPage ? data : [...prev, ...data]))
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
          setIsLoadingMore(false)
        }
        inFlightRef.current = false
      }
    }

    loadPokemons()

    return () => {
      isMounted = false
      controller.abort()
    }
  }, [page, pageSize])

  return {
    pokemons,
    isLoading,
    isLoadingMore,
    error,
    loadMore,
    hasMore,
  }
}

export default useInfinitePokemons
