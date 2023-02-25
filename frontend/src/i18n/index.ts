import en from './en'
import zh from './zh'
import ja from './ja'

export const Language = {
    en, zh, ja
}

export type Lang = keyof typeof Language

// i18n 结构
export type LocalizedType = typeof Language.zh

export function getDefaultLanguage(): Lang {
    for (const language of window.navigator.languages) {
        for (const lang of Object.keys(Language)) {
            if (language.includes(lang)) {
                return lang as Lang
            }
        }
    }
    return 'zh'
}
