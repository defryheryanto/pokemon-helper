import './App.css'
import { Button, Card, TypeBadge } from './components'

function App() {
  return (
    <div className="app">
      <header className="hero">
        <div>
          <p className="eyebrow">Pokemon Helper UI Kit</p>
          <h1>Neutral-first components with type accents.</h1>
          <p className="lede">
            Base UI stays clean and reusable. Type colors only appear where they
            should.
          </p>
        </div>
        <div className="hero-actions">
          <Button>Primary Action</Button>
          <Button variant="secondary">Secondary</Button>
          <Button variant="ghost">Ghost</Button>
          <Button variant="destructive">Destructive</Button>
        </div>
      </header>

      <section className="grid">
        <Card
          title="Team Builder"
          subtitle="Build squads with neutral UI elements."
        >
          <p>
            Cards stay neutral so accent colors can be reserved for Pokemon
            types.
          </p>
          <div className="badge-row">
            <TypeBadge type="fire" />
            <TypeBadge type="water" />
            <TypeBadge type="grass" />
          </div>
        </Card>

        <Card title="Strategy Notes" subtitle="Keep notes tidy and focused.">
          <p>
            Components use shared spacing, radius, and shadows for a consistent
            feel.
          </p>
          <div className="cta-row">
            <Button variant="secondary">Save Draft</Button>
            <Button>Publish</Button>
          </div>
        </Card>
      </section>
    </div>
  )
}

export default App
