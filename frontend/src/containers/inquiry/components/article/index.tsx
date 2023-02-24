import Empty from "@components/empty";
import {articleAtom, ArticleType} from "@stores/jotai";
import {useClient, useI18n, useVisible} from "@lib/hook";
import {Icon} from '@iconify/react';
import unfoldMoreHorizontal from '@iconify/icons-mdi/unfold-more-horizontal';
import unfoldLessHorizontal from '@iconify/icons-mdi/unfold-less-horizontal';
import './style.css'
import closeIcon from '@iconify/icons-mdi/close';
import classnames from "classnames";
import React, {
    CSSProperties,
    forwardRef,
    memo,
    Ref,
    useCallback,
    useEffect,
    useLayoutEffect,
    useRef,
    useState
} from "react";
import '@styles/markdown.css'
import {useAtomValue, useSetAtom} from "jotai";
import {Updater, useImmer} from "use-immer";
import {useArticleReader} from "@lib/request";
import {QAItem, QASubmissionBox} from "@containers/inquiry/components/article/qa";
import {noop} from "@lib/helper";

type ArticleItemHeader = Omit<ArticleType, 'records'> & {
    toggleMore: () => void
    unfoldMore: boolean
    expandText: boolean
    toggleExpandText: () => void
}

const ArticleItemHeader = memo(forwardRef((props: ArticleItemHeader, ref: Ref<HTMLDivElement>) => {
    const {
        title,
        content,
        createTime,
        unfoldMore,
        toggleMore,
        expandText,
        toggleExpandText
    } = props

    const {translation} = useI18n()
    const setRecords = useSetAtom(articleAtom)
    const clickClose = useCallback(() => {
        setRecords(draft => {
            return draft.filter(item => item.createTime !== createTime)
        })
    }, [setRecords])

    return <div
        flex='~ row'
        ref={ref}
        select='none'
    >
        <div className='article-item-icon' m='r-2 t-[13px]' onClick={
            () => {
                // 如果 expandText 为 true 时关闭，需要先关闭 expandText
                expandText && unfoldMore && toggleExpandText()
                toggleMore()
            }}>
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
            >{title}
            </p>
            <p text='xs md:sm space-pre-line'
               select='text'
               transition='colors'
               className={classnames('break-words',
                   {'line-clamp-1': !unfoldMore},
                   {'line-clamp-4': unfoldMore && !expandText},
                   {'text-gray-600': expandText},
                   {'text-slate-400': !expandText}
               )}
            >{content}</p>
            <div
                text='xs md:sm sky-800/80 hover:sky-800'
                font='medium hover:bold'
                cursor='pointer'
                m='t-2'
                className={classnames({"hidden": !unfoldMore})}
                onClick={toggleExpandText}
            >
                {translation("article.itemFrom.expendText")[expandText ? 1 : 0]}
            </div>
        </div>
        <div className='article-item-icon' m='t-[13px]' onClick={clickClose}>
            <Icon icon={closeIcon}/>
        </div>
    </div>
}))


interface ArticleItemBody extends Omit<ArticleType, 'records'> {
    records: ArticleType['records']
    setRecords: Updater<ArticleType['records']>
    createTime: ArticleType['createTime']
}

const ArticleItemBody = memo(forwardRef((props: ArticleItemBody, ref: Ref<HTMLDivElement>) => {
    const {
        records,
        setRecords,
        createTime,
        title,
        content
    } = props

    const setReadRecords = useSetAtom(articleAtom)

    // 清空上下文方法
    const clearContext = useCallback(() => {
        setReadRecords(records => {
            const item = records.find(item => item.createTime === createTime);
            if (!item) return records;
            item.records = [];
            return [...records]
        })
        setRecords(draft => {
            // 清空数组
            draft.splice(0, draft.length)
        })
    }, [setReadRecords, setRecords])

    const client = useClient();

    const childRef = useRef<{
        question: string,
        setErrorObj: (value: { msg: string }) => void,
        setIsLoad: (value: boolean) => void,
        setQuestion: (value: string) => void
    }>({
        question: "",
        setErrorObj: noop,
        setIsLoad: noop,
        setQuestion: noop
    })

    const onAnswerEnd = useCallback((endText: string, err: boolean) => {
        const {
            question,
            setErrorObj,
            setIsLoad,
            setQuestion
        } = childRef.current

        if (err) {
            setRecords(draft => {
                draft.splice(draft.length - 1, 1)
            })
            setErrorObj({msg: endText})
        } else setReadRecords(up => {
            const item = up.find(item => item.createTime === createTime);
            if (!item) return up;
            item.records = [...records.slice(0, records.length - 1), {Q: question, A: endText}]
            return [...up]
        })
        setQuestion("")
        setIsLoad(false)
    }, [setReadRecords, setRecords, records, createTime, childRef])

    // 待获取数据的ArticleId
    const [articleId, setArticleId] = useState<string | null>(null)
    const localeAnswer = useArticleReader(articleId, onAnswerEnd)
    useEffect(() => {
        if (!localeAnswer) return
        setRecords(draft => {
            // draft的最后一个
            const lastItem = draft[draft.length - 1]
            if (!lastItem) return
            lastItem.A = localeAnswer
        })
    }, [localeAnswer, setRecords])

    // 提交问题
    const submitQuestion = useCallback(async () => {
        const {
            question,
            setErrorObj,
            setIsLoad
        } = childRef.current

        setIsLoad(true)
        const res = await client.registerTopic("article", {
            noun: `${title} || ${question}`,
            description: content,
        });
        if (res.code === 200) {
            setRecords(draft => {
                draft.push({Q: question, A: ""})
            })
            setArticleId(res.data!.id)
        } else {
            setErrorObj({msg: res.message})
            setIsLoad(false)
        }
    }, [client, setRecords, setArticleId, childRef])

    return <div
        className='article-item-body' flex='~ col' space='y-2' ref={ref}>
        {
            records.map((item, index) => {
                return <QAItem {...item} key={index}/>
            })
        }
        <QASubmissionBox
            childRef={childRef}
            submitQuestion={submitQuestion}
            clearContext={clearContext}
        />
    </div>
}))


const ArticleItem = memo((props: ArticleType) => {
    const {
        records, createTime
    } = props

    const {
        visible: unfoldMore,
        toggle: toggleMore
    } = useVisible(false)
    const {visible: expandText, toggle: toggleExpandText} = useVisible(false)

    const [temporaryRecords, setTemporaryRecords] = useImmer(records)

    const headerRef = useRef<HTMLDivElement>(null)
    const bodyRef = useRef<HTMLDivElement>(null)

    const [itemStyle, setItemStyle] = useState<CSSProperties>(
        {
            height: window.innerWidth > 768 ? 80 : 76,
            transitionDuration: '0.02s'
        })
    useLayoutEffect(() => {
        const headerHeight = headerRef.current?.clientHeight!!
        const bodyHeight = bodyRef.current?.clientHeight ?? 0
        const height = headerHeight + bodyHeight + 32
        // perf : 如果 height 没有变化，那么就不要重新渲染
        if (itemStyle.height === height) return
        setItemStyle({
            height: height, transitionDuration: height / 4000 >= 0.2 ? `${height / 4000}s` : `0.2s`
        })
    }, [unfoldMore, temporaryRecords, expandText])

    return <div
        className='article-item'
        shadow='~'
        border='rounded'
        p='4'
        bg='white'
        style={itemStyle}
        overflow='clip'
    >
        <ArticleItemHeader
            ref={headerRef}
            unfoldMore={unfoldMore}
            toggleMore={toggleMore}
            expandText={expandText}
            toggleExpandText={toggleExpandText}
            {...props}
        />
        {unfoldMore && <ArticleItemBody
            {...props}
            ref={bodyRef}
            records={temporaryRecords}
            setRecords={setTemporaryRecords}
            createTime={createTime}
        />}
    </div>
})

const Article = () => {

    const articleRecord = useAtomValue(articleAtom)
    const {translation} = useI18n()

    return <div flex='1' m='y-2.5' space='y-2.5'>
        {
            articleRecord.length === 0 ?
                <Empty title={translation('empty')}/> :
                <>
                    {
                        articleRecord
                            .map((item) => {
                                return <ArticleItem key={item.createTime} {...item}/>
                            })
                    }
                </>
        }
    </div>
}
export default Article
