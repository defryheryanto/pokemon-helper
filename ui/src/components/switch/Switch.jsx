import styles from './Switch.module.css'

function Switch({ label, className = '', ...props }) {
  return (
    <label className={`${styles.wrapper} ${className}`.trim()}>
      <input className={styles.input} type="checkbox" {...props} />
      <span className={styles.track} aria-hidden="true">
        <span className={styles.thumb} />
      </span>
      {label ? <span className={styles.label}>{label}</span> : null}
    </label>
  )
}

export default Switch
