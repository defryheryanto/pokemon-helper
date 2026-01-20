import { useEffect } from 'react'
import { createPortal } from 'react-dom'
import styles from './Modal.module.css'

function Modal({ open, onClose, title, children, actions }) {
  useEffect(() => {
    if (!open) return

    const handleKey = (event) => {
      if (event.key === 'Escape') {
        onClose?.()
      }
    }

    document.addEventListener('keydown', handleKey)
    return () => document.removeEventListener('keydown', handleKey)
  }, [open, onClose])

  if (!open) return null

  return createPortal(
    <div className={styles.overlay} role="presentation" onClick={onClose}>
      <div
        className={styles.dialog}
        role="dialog"
        aria-modal="true"
        aria-label={title}
        onClick={(event) => event.stopPropagation()}
      >
        <div className={styles.header}>
          <h2 className={styles.title}>{title}</h2>
          <button type="button" className={styles.close} onClick={onClose}>
            Close
          </button>
        </div>
        <div className={styles.body}>{children}</div>
        {actions ? <div className={styles.footer}>{actions}</div> : null}
      </div>
    </div>,
    document.body,
  )
}

export default Modal
