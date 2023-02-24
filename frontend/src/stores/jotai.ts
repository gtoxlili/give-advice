import {atomWithStorage} from "jotai/utils";
import {getDefaultLanguage, Lang} from "@i18n";
import {atom} from "jotai";
import {Client} from "@lib/axios";

// 全局 i18n
export const localeAtom = atomWithStorage<Lang>('locale',
    getDefaultLanguage()
)

// 查询记录
// 主题 | 描述 | 答案 | 询问时间
export interface AdviceType {
    topic: string
    description: string
    answer: string
    askTime: number
}

export const adviceAtom = atomWithStorage<AdviceType[]>('advice-records', [])

// 文章记录
// 文章标题 | 文章内容 | 建立时间
export interface ArticleType {
    title: string
    content: string
    createTime: number
    records: {
        Q: string
        A: string
    }[]
}

export const articleAtom = atomWithStorage<ArticleType[]>('article-records', [])

// axios 实例
export const clientAtom = atom({
    key: '',
    instance: null as Client | null,
})

// 用户数据
interface User {
    auth: string
    url: string
    token: string
}

export const userAtom = atomWithStorage<User>('user-info', {
    url: '/api',
    auth: '',
    token: '',
})
