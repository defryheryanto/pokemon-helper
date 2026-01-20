import { typeColorMap } from '../theme'
import styles from './TypeBadge.module.css'

function TypeBadge({ type, className = '' }) {
  const tone = type?.toLowerCase()
  const badgeColor = typeColorMap[tone] || typeColorMap.dark
  const label = type ? type.charAt(0).toUpperCase() + type.slice(1) : 'Unknown'

  return (
    <span
      className={`${styles.badge} ${className}`.trim()}
      style={{ '--type-color': badgeColor }}
    >
      {label}
    </span>
  )
}

export default TypeBadge
