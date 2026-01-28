import { Link } from 'react-router-dom'
import { Button, ButtonLink } from '../index'
import iconText from '../../assets/icon-text.png'

function Navbar() {
  return (
    <nav className="topbar">
      <Link className="brand" to="/">
        <img src={iconText} alt="PokeLab" />
      </Link>
      <div className="nav-links">
        <ButtonLink variant="ghost" className="nav-link" to="/pokemons">
          Pokemon
        </ButtonLink>
        <ButtonLink variant="ghost" className="nav-link" to="/team-builder">
          Team Builder
        </ButtonLink>
      </div>
    </nav>
  )
}

export default Navbar
