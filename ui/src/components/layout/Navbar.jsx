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
        <ButtonLink variant="ghost" to="/pokemons">
          Pokemon
        </ButtonLink>
        <Button variant="ghost" type="button">
          Team Builder
        </Button>
      </div>
    </nav>
  )
}

export default Navbar
