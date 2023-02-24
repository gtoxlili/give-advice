import React, {memo, MutableRefObject, useEffect, useState} from "react";
import {ArticleType} from "@stores/jotai";
import {useImmer} from "use-immer";
import Textarea from "@components/textarea";
import Button from "@components/button";
import {useI18n} from "@lib/hook";
import ErrorTip from "@components/errorTip";
import "./style.css"

export const QAItem = memo((props: ArticleType['records'][0]) => {
    const {Q, A} = props
    return <div text='sm color-[#24292f]'>
        <div m="b-1" font='medium'>
            <span text='fuchsia-800' m='r-2' font='bold'>Q :</span>{Q}
        </div>
        <div text="color-[#262626]">
            <span>
                <span text='sky-700' m='r-2' font='bold'>A :</span>
            </span>
            <span text='xs md:sm justify'>{A}</span>
        </div>
    </div>
})

interface QASubmissionBoxProps {
    childRef: MutableRefObject<{
        question: string,
        setErrorObj: (value: { msg: string }) => void,
        setIsLoad: (value: boolean) => void,
        setQuestion: (value: string) => void
    }>,
    submitQuestion: () => Promise<void>,
    clearContext: () => void
}

export const QASubmissionBox = memo((props: QASubmissionBoxProps) => {

    const {childRef, submitQuestion, clearContext} = props

    const [question, setQuestion] = useImmer('')
    const [isLoad, setIsLoad] = useState(false)
    const {translation} = useI18n()
    const [errorObj, setErrorObj] = useState({msg: ""})

    useEffect(() => {
        childRef.current.question = question
        childRef.current.setErrorObj = setErrorObj
        childRef.current.setIsLoad = setIsLoad
        childRef.current.setQuestion = setQuestion
    }, [question, setErrorObj, setIsLoad, setQuestion])

    return <>
        <ErrorTip obj={errorObj}/>
        <Textarea
            value={question}
            rows={4}
            onChange={setQuestion}
            placeholder={translation("article.itemFrom.placeholder")}
            text='placeholder-sky-700/30 placeholder-italic'
            m='b-2'
            className="question-textarea"
        />
        <div flex='~ row-reverse' justify='between'>
            <Button
                loading={isLoad}
                fullWidth={window.innerWidth < 768}
                variant='light'
                color='#a21caf'
                onClick={submitQuestion}
            >{translation("article.itemFrom.submit")}
            </Button>
            {
                window.innerWidth >= 768 &&
                <Button
                    fullWidth={window.innerWidth < 768}
                    variant='subtle'
                    color='#be123c'
                    onClick={clearContext}
                >{translation("article.itemFrom.clear")}
                </Button>
            }
        </div>
    </>
})
