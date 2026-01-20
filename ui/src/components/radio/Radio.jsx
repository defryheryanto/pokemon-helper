import styles from './Radio.module.css'

function Radio({ label, className = '', ...props }) {
  return (
    <label className={`${styles.wrapper} ${className}`.trim()}>
      <input className={styles.input} type="radio" {...props} />
      <span className={styles.dot} aria-hidden="true" />
      <span className={styles.label}>{label}</span>
    </label>
  )
}

export default Radio
