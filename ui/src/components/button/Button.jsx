import styles from './Button.module.css'

const buttonVariants = {
  primary: styles.primary,
  secondary: styles.secondary,
  ghost: styles.ghost,
  destructive: styles.destructive,
}

function Button({ variant = 'primary', className = '', type = 'button', ...props }) {
  const variantClass = buttonVariants[variant] || buttonVariants.primary
  const classes = `${styles.button} ${variantClass} ${className}`.trim()

  return <button type={type} className={classes} {...props} />
}

export default Button
