import './App.css'
import { Button, Card, TypeBadge } from './components'

function App() {
  return (
    <div className="page">
      <div className="app">
        <nav className="topbar">
          <div className="brand">
            <span className="brand-dot" aria-hidden="true" />
            PokeLab
          </div>
          <div className="nav-links">
            <Button variant="ghost">Pokemon</Button>
            <Button variant="ghost">Team Builder</Button>
          </div>
        </nav>

        <header className="hero">
          <div className="hero-copy">
            <p className="eyebrow">Free forever</p>
            <h1>Build balanced Pokemon teams ready for any battle.</h1>
            <p className="lede">
              PokeLab helps you explore Pokemon data quickly and build a solid
              squad. Add your favorites, then the system recommends the Pokemon
              and elements you still need. No subscriptions, no paywalls.
            </p>
            <div className="hero-actions">
              <Button>Start Exploring</Button>
              <Button variant="secondary">See Features</Button>
            </div>
            <div className="hero-stats">
              <div className="stat">
                <span className="stat-number">800+</span>
                <span className="stat-label">Pokemon data</span>
              </div>
              <div className="stat">
                <span className="stat-number">18</span>
                <span className="stat-label">Element types</span>
              </div>
              <div className="stat">
                <span className="stat-number">100% free</span>
                <span className="stat-label">No limits</span>
              </div>
            </div>
          </div>

          <Card
            className="hero-card"
            title="Team Snapshot"
            subtitle="Fill 6 slots to highlight what your team needs."
          >
            <div className="team-preview">
              <div className="team-slot">Slot 1</div>
              <div className="team-slot">Slot 2</div>
              <div className="team-slot">Slot 3</div>
              <div className="team-slot muted">Slot 4</div>
              <div className="team-slot muted">Slot 5</div>
              <div className="team-slot muted">Slot 6</div>
            </div>
            <div className="badge-row">
              <TypeBadge type="fire" />
              <TypeBadge type="water" />
              <TypeBadge type="electric" />
              <TypeBadge type="grass" />
              <TypeBadge type="dragon" />
            </div>
            <p className="helper-text">
              Focus on coverage and roles so your team adapts to any matchup.
            </p>
          </Card>
        </header>

        <section className="section">
          <div className="section-header">
            <p className="eyebrow">Core features</p>
            <h2>Everything a trainer needs in one dashboard.</h2>
          </div>
          <div className="grid">
            <Card title="Pokemon Dex" subtitle="Find stats fast.">
              <p>
                Browse Pokemon lists, abilities, and elements in a structured
                way so you can pick a strong team core. All access is free.
              </p>
              <div className="cta-row">
                <Button variant="ghost">Preview data</Button>
              </div>
            </Card>

            <Card title="Team Builder" subtitle="Build teams with insights.">
              <p>
                Choose your favorites, then instantly review roles, weaknesses,
                and type coverage. Build as many teams as you want.
              </p>
              <div className="cta-row">
                <Button variant="ghost">Build a team</Button>
              </div>
            </Card>

            <Card title="Element Suggestions" subtitle="System calculated strategy tips.">
              <p>
                The system suggests Pokemon, types, and elements that are still
                missing so your composition gets stronger, at no cost.
              </p>
              <div className="badge-row">
                <TypeBadge type="ice" />
                <TypeBadge type="psychic" />
                <TypeBadge type="dark" />
              </div>
            </Card>
          </div>
        </section>

        <section className="section split-section">
          <div className="section-header">
            <p className="eyebrow">Workflow</p>
            <h2>Fast steps to build your best squad.</h2>
            <p className="lede">
              The flow stays short so you can focus on strategy and synergy.
            </p>
          </div>
          <div className="steps">
            <Card title="1. Explore data" subtitle="Find Pokemon and types.">
              <p>
                Use the Pokemon menu to check stats, types, and strengths for
                each character.
              </p>
            </Card>
            <Card title="2. Build a team" subtitle="Choose 3-6 Pokemon.">
              <p>
                Add your core members, then review coverage and roles on the
                dashboard.
              </p>
            </Card>
            <Card title="3. Get recommendations" subtitle="App helps you finish.">
              <p>
                The system suggests Pokemon or elements that are missing so your
                team stays balanced.
              </p>
            </Card>
          </div>
        </section>

        <section className="cta-section">
          <Card
            className="cta-card"
            title="Ready to become a top trainer?"
            subtitle="Start with clean data and the right recommendations for free."
          >
            <p>
              PokeLab is ready to guide your journey. Check data or start
              building your team right now. No subscriptions required.
            </p>
            <div className="cta-row">
              <Button>Explore Pokemon</Button>
              <Button variant="secondary">Build Team</Button>
            </div>
          </Card>
        </section>

        <footer className="footer">
          <span className="powered">
            Powered by{' '}
            <a href="https://pokeapi.co/" target="_blank" rel="noreferrer">
              PokeAPI
            </a>
          </span>
        </footer>
      </div>
    </div>
  )
}

export default App
