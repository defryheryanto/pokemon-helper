import styles from './Button.module.css'

const buttonVariants = {
  primary: styles.primary,
  secondary: styles.secondary,
  ghost: styles.ghost,
  destructive: styles.destructive,
}

export function getButtonClassName(variant = 'primary', className = '') {
  const variantClass = buttonVariants[variant] || buttonVariants.primary
  return `${styles.button} ${variantClass} ${className}`.trim()
}
