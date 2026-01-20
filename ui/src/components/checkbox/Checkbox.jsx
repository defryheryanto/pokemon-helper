import styles from './Checkbox.module.css'

function Checkbox({ label, className = '', ...props }) {
  return (
    <label className={`${styles.wrapper} ${className}`.trim()}>
      <input className={styles.input} type="checkbox" {...props} />
      <span className={styles.box} aria-hidden="true" />
      <span className={styles.label}>{label}</span>
    </label>
  )
}

export default Checkbox
