import {Select, SelectOptions} from "@components/select";
import React, {useCallback, useState} from "react";
import {useI18n} from "@lib/hook";
import {Lang} from "@i18n";
import {userAtom} from "@stores/jotai";
import Input from "@components/input";
import {useImmerAtom} from "jotai-immer";
import "./style.css"

const languageOptions: SelectOptions<Lang>[] = [
    {label: '中文', value: 'zh'},
    {label: 'English', value: 'en'},
    {label: '日本語', value: 'ja'}
]

// Github Svg
export const GithubSvg = () => {
    return <svg aria-hidden="true" height="20" viewBox="0 0 16 16" version="1.1" width="24" data-view-component="true"
                className="github-icon">
        <path fillRule="evenodd"
              d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"></path>
    </svg>
}

export const LanguageSelect = () => {
    const {locale, setLocale, translation} = useI18n()

    return <div
        flex='~ row'
        align='items-center'
        justify='between'
        text="color-[#54759a] sm"
        font='semibold'
    >
        {translation("settings.first.language")}
        <Select options={languageOptions} onSelect={setLocale} value={locale}/>
    </div>
}

// 自备 Token
export const OwnToken = () => {
    const [{token}, setToken] = useImmerAtom(userAtom)

    const [temporaryToken, setTemporaryToken] = useState(token)
    const {translation} = useI18n()

    const onBlur = useCallback(
        () => {
            setToken(draft => {
                draft.token = temporaryToken
            })
        }, [temporaryToken, setToken])

    return <div
        flex='~ row'
        align='items-center'
        justify='between'
        text="color-[#54759a] sm"
        font='semibold'
    >
        {translation("settings.first.ownToken")}
        <div w='40'>
            <Input value={temporaryToken} onChange={setTemporaryToken} onBlur={onBlur}/>
        </div>
    </div>
}

// 项目链接
export const ProjectLink = () => {
    const {translation} = useI18n()

    return <div
        flex='~ row'
        align='items-center'
        justify='between'
        text="color-[#54759a] sm"
        font='semibold'
    >
        {translation("settings.second.project")}
        <a href="https://github.com/gtoxlili/advice-hub"
           target="_blank"
           rel="noreferrer"
           className='text-slate-500 hover:text-sky-700' title='Github'>
            <GithubSvg/>
        </a>
    </div>
}
