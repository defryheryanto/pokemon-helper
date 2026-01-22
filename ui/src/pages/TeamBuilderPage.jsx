import { useEffect, useMemo, useRef, useState } from 'react'
import { Button, Card, TypeBadge } from '../components'
import useInfinitePokemons from '../hooks/useInfinitePokemons'
import { fetchTeamSimulation } from '../services/teamBuilderService'

function TeamBuilderPage() {
  const {
    pokemons,
    isLoading,
    isLoadingMore,
    error,
    loadMore,
    hasMore,
  } = useInfinitePokemons({ pageSize: 36 })
  const [team, setTeam] = useState([])
  const [coveredTypes, setCoveredTypes] = useState([])
  const [uncoveredTypes, setUncoveredTypes] = useState([])
  const [suggestionTypes, setSuggestionTypes] = useState([])
  const [isSimulating, setIsSimulating] = useState(false)
  const [simulateError, setSimulateError] = useState(null)
  const listRef = useRef(null)
  const sentinelRef = useRef(null)

  useEffect(() => {
    const sentinel = sentinelRef.current
    const list = listRef.current
    if (!sentinel || !list || !hasMore) {
      return undefined
    }

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries[0]?.isIntersecting) {
          loadMore()
        }
      },
      { root: list, rootMargin: '160px' },
    )

    observer.observe(sentinel)

    return () => observer.disconnect()
  }, [hasMore, loadMore])

  useEffect(() => {
    if (team.length === 0) {
      setCoveredTypes([])
      setUncoveredTypes([])
      setSuggestionTypes([])
      setSimulateError(null)
      return undefined
    }

    const controller = new AbortController()
    let isMounted = true

    async function simulateTeam() {
      setIsSimulating(true)
      setSimulateError(null)
      try {
        const payload = await fetchTeamSimulation(
          {
            pokemons: team.map((pokemon) => pokemon.name),
            with_type_suggestion: true,
          },
          { signal: controller.signal },
        )

        if (!isMounted) {
          return
        }

        const normalizeTypes = (types = []) => {
          const seen = new Set()
          return types.filter((type) => {
            if (!type) {
              return false
            }
            const key = type.toLowerCase()
            if (seen.has(key)) {
              return false
            }
            seen.add(key)
            return true
          })
        }

        setCoveredTypes(normalizeTypes(payload?.covered_types))
        setUncoveredTypes(normalizeTypes(payload?.uncovered_types))
        setSuggestionTypes(normalizeTypes(payload?.suggestion_types))
      } catch (err) {
        if (err.name === 'AbortError') {
          return
        }
        if (isMounted) {
          setSimulateError(err)
        }
      } finally {
        if (isMounted) {
          setIsSimulating(false)
        }
      }
    }

    simulateTeam()

    return () => {
      isMounted = false
      controller.abort()
    }
  }, [team])

  const teamSlots = useMemo(
    () => Array.from({ length: 6 }, (_, index) => team[index] ?? null),
    [team],
  )

  const teamNames = useMemo(() => new Set(team.map((pokemon) => pokemon.name)), [team])

  const handleAdd = (pokemon) => {
    setTeam((prev) => {
      if (prev.length >= 6 || prev.some((member) => member.name === pokemon.name)) {
        return prev
      }
      return [...prev, pokemon]
    })
  }

  const handleRemove = (pokemon) => {
    setTeam((prev) => prev.filter((member) => member.name !== pokemon.name))
  }

  return (
    <section className="team-builder-page">
      <div className="section-header">
        <p className="eyebrow">Team builder</p>
        <h2>Draft a squad and check your coverage.</h2>
        <p className="lede">
          Pick up to six Pokemon, then review the types your team can hit.
        </p>
      </div>

      <div className="team-builder-layout">
        <Card
          className="team-list-card"
          title="Pokemon list"
          subtitle="Add up to six Pokemon to simulate coverage."
        >
          <div className="team-list-body">
            {isLoading ? <div className="pokemon-status">Loading Pokemon...</div> : null}
            {error ? (
              <div className="pokemon-status error">
                Sorry, we could not load the Pokemon list.
              </div>
            ) : null}

            {!isLoading && !error ? (
              <div className="team-list" ref={listRef}>
                {pokemons.map((pokemon) => {
                  const isSelected = teamNames.has(pokemon.name)
                  return (
                    <div className="team-list-item" key={`${pokemon.id}-${pokemon.name}`}>
                      <div className="team-list-sprite">
                        <img
                          src={pokemon.sprites}
                          alt={pokemon.name}
                          loading="lazy"
                        />
                      </div>
                      <div className="team-list-info">
                        <span className="team-list-name">{pokemon.name}</span>
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
                      <Button
                        variant={isSelected ? 'secondary' : 'ghost'}
                        className="team-list-action"
                        onClick={() =>
                          isSelected ? handleRemove(pokemon) : handleAdd(pokemon)
                        }
                        disabled={!isSelected && team.length >= 6}
                      >
                        {isSelected ? 'Remove' : 'Add'}
                      </Button>
                    </div>
                  )
                })}
                {hasMore ? (
                  <div ref={sentinelRef} className="team-list-sentinel" />
                ) : null}
              </div>
            ) : null}

            {!error && isLoadingMore ? (
              <div className="pokemon-status">Loading more Pokemon...</div>
            ) : null}
          </div>
        </Card>

        <Card
          className="team-summary-card"
          title="Selected team"
          subtitle="Slots auto-fill based on your picks."
        >
          <div className="team-preview team-preview--builder">
            {teamSlots.map((pokemon, index) => (
              <button
                type="button"
                key={`team-slot-${index}`}
                className={`team-slot team-slot-button ${pokemon ? 'filled' : 'muted'}`}
                onClick={() => (pokemon ? handleRemove(pokemon) : null)}
                disabled={!pokemon}
              >
                {pokemon ? (
                  <div className="team-slot-content">
                    <div className="team-slot-sprite">
                      <img src={pokemon.sprites} alt={pokemon.name} loading="lazy" />
                    </div>
                    <span className="team-slot-name">{pokemon.name}</span>
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
                ) : (
                  `Slot ${index + 1}`
                )}
              </button>
            ))}
          </div>

          <div className="team-coverage">
            <div className="team-coverage-section">
              <p className="team-coverage-title">Type coverage</p>
              <div className="badge-row">
                {coveredTypes.length ? (
                  coveredTypes.map((type) => <TypeBadge key={type} type={type} />)
                ) : (
                  <span className="helper-text">Add Pokemon to see coverage.</span>
                )}
              </div>
            </div>
            <div className="team-coverage-section">
              <p className="team-coverage-title">Types not covered</p>
              <div className="badge-row">
                {uncoveredTypes.length ? (
                  uncoveredTypes.map((type) => <TypeBadge key={type} type={type} />)
                ) : (
                  <span className="helper-text">No gaps yet.</span>
                )}
              </div>
            </div>
            <div className="team-coverage-section">
              <p className="team-coverage-title">Type suggestion</p>
              <div className="badge-row">
                {suggestionTypes.length ? (
                  suggestionTypes.map((type) => <TypeBadge key={type} type={type} />)
                ) : (
                  <span className="helper-text">Add Pokemon to get suggestions.</span>
                )}
              </div>
            </div>
          </div>

          {isSimulating ? (
            <p className="helper-text">Simulating team coverage...</p>
          ) : null}
          {simulateError ? (
            <p className="helper-text team-coverage-error">
              Could not load team coverage. Try again.
            </p>
          ) : null}
        </Card>
      </div>
    </section>
  )
}

export default TeamBuilderPage
