import styles from './Badge.module.css'

function Badge({ children, className = '' }) {
  return <span className={`${styles.badge} ${className}`.trim()}>{children}</span>
}

export default Badge
