import './App.css'
import { Routes, Route } from 'react-router-dom'
import Navbar from './components/layout/Navbar'
import HomePage from './pages/HomePage'
import PokemonsPage from './pages/PokemonsPage'
import TeamBuilderPage from './pages/TeamBuilderPage'

function App() {
  return (
    <div className="page">
      <div className="app">
        <Navbar />

        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/pokemons" element={<PokemonsPage />} />
          <Route path="/team-builder" element={<TeamBuilderPage />} />
        </Routes>

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
