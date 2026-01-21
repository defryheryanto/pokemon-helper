import { Link } from 'react-router-dom'
import { getButtonClassName } from './buttonStyles'

function ButtonLink({ variant = 'primary', className = '', to, ...props }) {
  const classes = getButtonClassName(variant, className)
  return <Link to={to} className={classes} {...props} />
}

export default ButtonLink
