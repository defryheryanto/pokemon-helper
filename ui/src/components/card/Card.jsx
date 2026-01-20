import styles from './Card.module.css'

function Card({ title, subtitle, children, className = '' }) {
  return (
    <div className={`${styles.card} ${className}`.trim()}>
      {title ? (
        <div className={styles.header}>
          <h3 className={styles.title}>{title}</h3>
          {subtitle ? <p className={styles.subtitle}>{subtitle}</p> : null}
        </div>
      ) : null}
      <div className={styles.body}>{children}</div>
    </div>
  )
}

export default Card
