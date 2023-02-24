import {useCallback, useDeferredValue, useEffect, useLayoutEffect, useMemo, useRef, useState} from "react";
import {AdviceType, clientAtom, localeAtom, userAtom} from "@stores/jotai";
import {useAtom, useAtomValue} from "jotai";
import {Language, LocalizedType} from "@i18n";
import {Get} from "type-fest";
import {get} from "lodash-es";
import Infer from "@lib/type";
import {useImmer} from "use-immer";
import {Draft} from "immer";
import {Client} from "@lib/axios";
import {marked} from "marked";

export function useI18n() {
    const [locale, setLocale] = useAtom(localeAtom)
    const translation = useCallback(<P extends Infer<LocalizedType>>(path: P) => {
        return get(Language[locale], path) as Get<LocalizedType, P>
    }, [locale])

    return {
        locale,
        setLocale,
        translation,
    }
}

export function useVisible(initial = false) {
    const [visible, setVisible] = useState(initial)
    const hide = useCallback(() => {
        setVisible(false)
    }, [])
    const show = useCallback(() => {
        setVisible(true)
    }, [])
    const toggle = useCallback(() => {
        setVisible((prev) => !prev)
    }, [])
    return {
        visible,
        hide,
        show,
        toggle,
    }
}

export function useObject<T extends Record<string, unknown>>(initialValue: T) {
    const [copy, rawSet] = useImmer(initialValue)

    const set = useMemo(<K extends keyof Draft<T>>() => {
        const set: Record<keyof Draft<T>, (value: Draft<T>[K]) => void> = Object.create(null)
        for (const key of Object.keys(initialValue)) {
            const k = key as keyof Draft<T>
            set[k] = (value: Draft<T>[K]) => {
                rawSet(draft => {
                    draft[k] = value
                })
            }
        }
        return set
    }, [rawSet])

    return [copy, set] as [T, typeof set]
}

export function useClient() {
    const userInfo = useAtomValue(userAtom)
    const [item, setItem] = useAtom(clientAtom)
    if (item.key !== userInfo.url) {
        setItem({
            key: userInfo.url,
            instance: new Client(userInfo.url, userInfo.auth,userInfo.token),
        })
    }
    return item.instance!
}

export function usePrevious<T>(value: T): T | undefined {
    const ref = useRef<T>();

    useEffect(() => {
        ref.current = value;
    }, [value]);

    return ref.current;
}
