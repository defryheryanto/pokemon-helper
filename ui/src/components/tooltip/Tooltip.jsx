import { useId } from 'react'
import styles from './Tooltip.module.css'

function Tooltip({ content, children, className = '' }) {
  const tooltipId = useId()

  return (
    <span className={`${styles.wrapper} ${className}`.trim()}>
      <span className={styles.trigger} aria-describedby={tooltipId} tabIndex={0}>
        {children}
      </span>
      <span role="tooltip" id={tooltipId} className={styles.tooltip}>
        {content}
      </span>
    </span>
  )
}

export default Tooltip
