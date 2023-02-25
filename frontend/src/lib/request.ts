import {useDeferredValue, useLayoutEffect, useMemo, useRef, useState} from "react";
import {marked} from "marked";

export function useInquiryReader(
    id: string | null,
    // 对于传回来的字符串如何处理
    onMessage: (message: string) => string,
    onEnd?: (endText: string, err: boolean) => void,
) {
    const [value, setValue] = useState<string>("")
    const valueRef = useRef(value)
    useLayoutEffect(() => {
        valueRef.current = value
    }, [value])

    const deferredText = useDeferredValue(value)
    const resultText = useMemo(() => {
        return onMessage(deferredText)
    }, [deferredText])

    useLayoutEffect(() => {
        if (!id) return
        setValue("")
        const source = new EventSource(`/api/inquiry/${id}`)
        source.onmessage = (event) => setValue(value => value + (event.data === '\\n' ? '\n' : event.data))
        // @ts-ignore
        source.onerror = (event: Event & { data: string }) => {
            if (event.data) {
                setValue(event.data)
                onEnd && onEnd(event.data, true)
            }
            source.close()
        }
        onEnd && source.addEventListener("end", () => onEnd(onMessage(valueRef.current), false))
        return () => source.close()
    }, [id])

    return resultText
}

export function useAdviceReader(id: string | null, onEnd?: (endText: string, err: boolean) => void) {
    return useInquiryReader(id,
        str => marked(str, {
            gfm: true,
            breaks: true,
            smartypants: true
        }), onEnd)
}

export function useArticleReader(id: string | null, onEnd?: (endText: string, err: boolean) => void) {
    return useInquiryReader(id,
        str => str
        , onEnd)
}
