import styles from './Toast.module.css'

const toastVariants = {
  success: styles.success,
  warning: styles.warning,
  error: styles.error,
  info: styles.info,
}

function Toast({ variant = 'info', title, children, className = '' }) {
  const variantClass = toastVariants[variant] || toastVariants.info
  const classes = `${styles.toast} ${variantClass} ${className}`.trim()

  return (
    <div className={classes} role="status" aria-live="polite">
      {title ? <div className={styles.title}>{title}</div> : null}
      {children ? <div className={styles.body}>{children}</div> : null}
    </div>
  )
}

export default Toast
