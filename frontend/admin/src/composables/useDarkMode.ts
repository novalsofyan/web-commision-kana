import { ref } from 'vue'

type Theme = 'light' | 'dark' | 'system'

const currentTheme = ref<Theme>('system')
const isDark = ref<boolean>(false)

export function useDarkMode() {
  const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

  const applyTheme = (theme: Theme): void => {
    let themeToApply: 'light' | 'dark'

    if (theme === 'system') {
      themeToApply = mediaQuery.matches ? 'dark' : 'light'
    } else {
      themeToApply = theme
    }

    document.documentElement.setAttribute('data-theme', themeToApply)
    isDark.value = themeToApply === 'dark'
    currentTheme.value = theme
  }

  const setTheme = (theme: Theme): void => {
    localStorage.setItem('theme-preference', theme)
    applyTheme(theme)
  }

  const initTheme = (): void => {
    const saved = localStorage.getItem('theme-preference') as Theme | null
    const initial = saved || 'system'
    applyTheme(initial)

    mediaQuery.addEventListener('change', () => {
      if (currentTheme.value === 'system') {
        applyTheme('system')
      }
    })
  }

  return { isDark, currentTheme, setTheme, initTheme }
}
