import { useEffect, useRef } from 'react'
import { TypeBadge } from '../components'
import useInfinitePokemons from '../hooks/useInfinitePokemons'

function PokemonsPage() {
  const {
    pokemons,
    isLoading,
    isLoadingMore,
    error,
    loadMore,
    hasMore,
  } = useInfinitePokemons({ pageSize: 60 })
  const sentinelRef = useRef(null)

  useEffect(() => {
    const sentinel = sentinelRef.current
    if (!sentinel || !hasMore) {
      return undefined
    }

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0]?.isIntersecting) {
          loadMore()
        }
      },
      { rootMargin: '200px' },
    )

    observer.observe(sentinel)

    return () => observer.disconnect()
  }, [hasMore, loadMore])

  return (
    <section className="pokemons-page">
      <div className="section-header">
        <p className="eyebrow">Pokemon index</p>
        <h2>Browse the full PokeLab roster.</h2>
        <p className="lede">
          Fresh data loaded straight from the PokeLab backend.
        </p>
      </div>

      {isLoading ? (
        <div className="pokemon-status">Loading Pokemon...</div>
      ) : null}
      {error ? (
        <div className="pokemon-status error">
          Sorry, we could not load the Pokemon list.
        </div>
      ) : null}

      {!isLoading && !error ? (
        <div className="pokemon-grid">
          {pokemons.map((pokemon) => (
            <div className="pokemon-card" key={`${pokemon.id}-${pokemon.name}`}>
              <div className="pokemon-sprite">
                <img src={pokemon.sprites} alt={pokemon.name} loading="lazy" />
              </div>
              <span className="pokemon-name">{pokemon.name}</span>
              {Array.isArray(pokemon.types) && pokemon.types.length ? (
                <div className="badge-row pokemon-type-row">
                  {pokemon.types.map((type, index) => (
                    <TypeBadge
                      key={`${pokemon.id}-${String(type)}-${index}`}
                      type={String(type)}
                    />
                  ))}
                </div>
              ) : null}
            </div>
          ))}
        </div>
      ) : null}

      {!error && hasMore ? <div ref={sentinelRef} className="pokemon-sentinel" /> : null}
      {isLoadingMore ? (
        <div className="pokemon-status">Loading more Pokemon...</div>
      ) : null}
    </section>
  )
}

export default PokemonsPage
