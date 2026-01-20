import { useId } from 'react'
import styles from './Select.module.css'

function Select({ label, helperText, error, className = '', id, children, ...props }) {
  const autoId = useId()
  const selectId = id || autoId
  const describedBy = helperText ? `${selectId}-help` : undefined
  const message = error || helperText

  return (
    <div className={`${styles.field} ${className}`.trim()}>
      {label ? (
        <label className={styles.label} htmlFor={selectId}>
          {label}
        </label>
      ) : null}
      <select
        id={selectId}
        className={styles.select}
        aria-invalid={Boolean(error)}
        aria-describedby={describedBy}
        {...props}
      >
        {children}
      </select>
      {message ? (
        <div
          id={describedBy}
          className={error ? styles.error : styles.helper}
        >
          {message}
        </div>
      ) : null}
    </div>
  )
}

export default Select
