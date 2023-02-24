import en from './en'
import zh from './zh'
import ja from './ja'
import Infer from "@lib/type";
import {IsEqual} from "type-fest";

export const Language = {
    en, zh, ja
}

export type Lang = keyof typeof Language

type US = typeof Language.en
type CN = typeof Language.zh


// i18n 结构
export type LocalizedType = CN

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
