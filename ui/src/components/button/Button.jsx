import { getButtonClassName } from './buttonStyles'

function Button({ variant = 'primary', className = '', type = 'button', ...props }) {
  const classes = getButtonClassName(variant, className)

  return <button type={type} className={classes} {...props} />
}

export default Button
