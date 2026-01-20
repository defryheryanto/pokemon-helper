import { useId, useState } from 'react'
import styles from './Tabs.module.css'

function Tabs({ tabs = [], defaultTabId, onTabChange, className = '' }) {
  const tabsId = useId()
  const initialId = defaultTabId || tabs[0]?.id
  const [activeId, setActiveId] = useState(initialId)

  const handleSelect = (tabId) => {
    setActiveId(tabId)
    if (onTabChange) {
      onTabChange(tabId)
    }
  }

  return (
    <div className={`${styles.tabs} ${className}`.trim()}>
      <div role="tablist" aria-label="Tabs" className={styles.tabList}>
        {tabs.map((tab) => {
          const isActive = tab.id === activeId
          const buttonId = `${tabsId}-${tab.id}-tab`
          const panelId = `${tabsId}-${tab.id}-panel`

          return (
            <button
              key={tab.id}
              id={buttonId}
              role="tab"
              type="button"
              aria-selected={isActive}
              aria-controls={panelId}
              tabIndex={isActive ? 0 : -1}
              className={`${styles.tab} ${isActive ? styles.tabActive : ''}`}
              onClick={() => handleSelect(tab.id)}
              disabled={tab.disabled}
            >
              {tab.label}
            </button>
          )
        })}
      </div>
      <div className={styles.panelWrap}>
        {tabs.map((tab) => {
          const isActive = tab.id === activeId
          const panelId = `${tabsId}-${tab.id}-panel`
          const buttonId = `${tabsId}-${tab.id}-tab`

          return (
            <div
              key={tab.id}
              id={panelId}
              role="tabpanel"
              aria-labelledby={buttonId}
              hidden={!isActive}
              className={styles.panel}
            >
              {tab.content}
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default Tabs
