import Empty from "@components/empty";
import {adviceAtom, AdviceType} from "@stores/jotai";
import {useI18n, useVisible} from "@lib/hook";
import {Icon} from '@iconify/react';
import unfoldMoreHorizontal from '@iconify/icons-mdi/unfold-more-horizontal';
import unfoldLessHorizontal from '@iconify/icons-mdi/unfold-less-horizontal';
import './style.css'
import closeIcon from '@iconify/icons-mdi/close';
import classnames from "classnames";
import {CSSProperties, forwardRef, memo, Ref, useCallback, useLayoutEffect, useMemo, useRef, useState} from "react";
import '@styles/markdown.css'
import {Updater} from "use-immer";
import {LoadedType} from "@containers/inquiry";
import {useAtom, useSetAtom} from "jotai";
import {noop} from "@lib/helper";
import {useAdviceReader} from "@lib/request";

// item内容
type AdviceItemProps = AdviceType &
    { initialOn?: boolean, reset: (endText: string, err: boolean) => void }

type AdviceItemHeader = Omit<AdviceItemProps, 'answer'> & {
    toggleMore: () => void
    unfoldMore: boolean
}

interface AdviceItemBody {
    innerHTML: { __html: string }
}

const AdviceItemHeader = memo(forwardRef((props: AdviceItemHeader, ref: Ref<HTMLDivElement>) => {
    const {
        topic,
        description,
        unfoldMore,
        toggleMore,
        askTime,
        reset
    } = props

    const setRecords = useSetAtom(adviceAtom)
    const clickClose = useCallback(() => {
        // @ts-ignore
        reset()
        setRecords(draft => {
            return draft.filter(item => item.askTime !== askTime)
        })
    }, [setRecords])


    return <div
        flex='~ row'
        ref={ref}
        select='none'
    >
        <div className='advice-item-icon' m='r-2 t-[13px]' onClick={toggleMore}>
            <Icon icon={
                unfoldMore ? unfoldLessHorizontal : unfoldMoreHorizontal
            }/>
        </div>
        <div
            m='r-2'
            flex='1'
        >
            <p text='base fuchsia-800 shadow-sm'
               font='medium tracking-wide'
               m='b-1'
            >{topic}
            </p>
            <p text='sm gray-400 space-pre-wrap'
               className={classnames('break-words', {'line-clamp-1': !unfoldMore})}
            >{description}</p>
        </div>
        <div className='advice-item-icon' m='t-[13px]' onClick={clickClose}>
            <Icon icon={closeIcon}/>
        </div>
    </div>
}))

const AdviceItemBody = memo(forwardRef((props: AdviceItemBody, ref: Ref<HTMLDivElement>) => {
    const {
        innerHTML
    } = props

    return <div
        text='xs md:sm'
        className={classnames('markdown-body', 'advice-item-body')}
        ref={ref}
        dangerouslySetInnerHTML={innerHTML}
    />
}))


const AdviceItem = memo((props: AdviceItemProps) => {
    const {
        topic,
        description,
        answer,
        initialOn = false,
        askTime,
        reset
    } = props

    const {
        visible: unfoldMore,
        toggle: toggleMore
    } = useVisible(initialOn)

    const headerRef = useRef<HTMLDivElement>(null)
    const bodyRef = useRef<HTMLDivElement>(null)

    const [itemStyle, setItemStyle] = useState<CSSProperties>(
        {
            height: 80,
            transitionDuration: '0.02s'
        })
    useLayoutEffect(() => {
        const headerHeight = headerRef.current?.clientHeight!!
        const bodyHeight = bodyRef.current?.clientHeight ?? 0
        const height = headerHeight + bodyHeight + 32
        // perf : 如果 height 没有变化，那么就不要重新渲染
        if (itemStyle.height === height) return
        setItemStyle({
            height: height, transitionDuration: unfoldMore ? `${height / 4000}s` : `0.2s`
        })
    }, [unfoldMore, answer])

    const innerHTML = useMemo(
        () => {
            return {
                __html: answer
            }
        }, [answer])

    return <div
        className='advice-item'
        shadow='~'
        border='rounded'
        p='4'
        bg='white'
        style={itemStyle}
        overflow='clip'
    >
        <AdviceItemHeader ref={headerRef} unfoldMore={unfoldMore}
                          toggleMore={toggleMore}
                          topic={topic}
                          askTime={askTime}
                          description={description} reset={reset}/>
        {unfoldMore && <AdviceItemBody ref={bodyRef} innerHTML={innerHTML}/>}
    </div>
})

interface LoadedItemProps {
    loadedItem: LoadedType,
    reset: (endText: string, err: boolean) => void,

}

const LoadedItem = memo((props: LoadedItemProps) => {
    const {loadedItem, reset} = props
    const answer = useAdviceReader(loadedItem.id, reset)
    return <AdviceItem {...loadedItem} answer={answer} reset={reset} initialOn/>
})

interface AdviceProps {
    loadedItem: LoadedType,
    setLoadedItem: Updater<LoadedType>
}


const Advice = (props: AdviceProps) => {

    const {loadedItem, setLoadedItem} = props

    const [record, setRecord] = useAtom(adviceAtom)
    const {locale, translation} = useI18n()

    const resetLoaded = useCallback((endText: string, err: boolean) => {
        setRecord(draft => {
            return err ? draft : [{
                ...loadedItem,
                answer: endText
            }, ...draft.filter((item) => item.askTime !== loadedItem.askTime)]
        })
        setLoadedItem(draft => {
            draft.id = null
            draft.error = err
        })
    }, [setLoadedItem, setRecord, loadedItem, locale])

    return <div flex='1' m='y-2.5' space='y-2.5'>
        {
            record.length === 0 && loadedItem.id === null && !loadedItem.error ?
                <Empty title={translation('empty')}/> :
                <>
                    {
                        (loadedItem.id !== null || loadedItem.error) &&
                        <LoadedItem loadedItem={loadedItem} reset={resetLoaded}/>
                    }
                    {
                        record
                            .map((item) => {
                                return <AdviceItem
                                    key={item.askTime} {...item}
                                    initialOn={item.askTime === loadedItem.askTime}
                                    reset={noop}/>
                            })
                    }
                </>
        }
    </div>
}
export default Advice
