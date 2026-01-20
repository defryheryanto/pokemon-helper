import { useId } from 'react'
import styles from './Input.module.css'

function Input({ label, helperText, error, className = '', id, ...props }) {
  const autoId = useId()
  const inputId = id || autoId
  const describedBy = helperText ? `${inputId}-help` : undefined
  const message = error || helperText

  return (
    <div className={`${styles.field} ${className}`.trim()}>
      {label ? (
        <label className={styles.label} htmlFor={inputId}>
          {label}
        </label>
      ) : null}
      <input
        id={inputId}
        className={styles.input}
        aria-invalid={Boolean(error)}
        aria-describedby={describedBy}
        {...props}
      />
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

export default Input
