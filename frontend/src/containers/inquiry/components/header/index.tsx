import React, {useCallback, useRef, useState} from 'react'
import {useClient, useI18n, useObject, useVisible} from "@lib/hook";
import Button from "@components/button";
import contentSavePlusOutline from "@iconify/icons-mdi/content-save-plus-outline";
import {Icon} from "@iconify/react";
import Dialog from "@components/dialog";
import Input from "@components/input";
import Textarea from "@components/textarea";
import dayjs from "dayjs";
import {Updater} from "use-immer";
import {LoadedType} from "@containers/inquiry";
import {InquiryType} from "@containers/app";
import {useSetAtom} from "jotai";
import {articleAtom} from "@stores/jotai";
import ErrorTip from "@components/errorTip";

const AddItemModal = (
    props: {
        hide: () => void,
        setLoadedItem: Updater<LoadedType>
        inquiryType: InquiryType
    }
) => {
    const {
        hide,
        setLoadedItem,
        inquiryType
    } = props

    const {translation} = useI18n()
    const client = useClient();

    const [form, set] = useObject({
        noun: '',
        description: ''
    })

    const [errorObj, setErrorObj] = useState({msg: ""})

    const inSubmit = useRef(false)

    const onAdviceSubmit = useCallback(async () => {
        if (inSubmit.current) return
        inSubmit.current = true
        const {noun, description} = form
        const res = await client.registerTopic("advice", {
            noun,
            description
        });
        if (res.code === 200) {
            setLoadedItem(draft => {
                draft.topic = noun
                draft.description = description
                draft.id = res.data!.id
                draft.askTime = dayjs().valueOf()
            })
            hide()
            inSubmit.current = false
        } else {
            setErrorObj({msg: res.message})
            inSubmit.current = false
        }
    }, [client, hide, form, setLoadedItem, setErrorObj])

    const setArticle = useSetAtom(articleAtom)
    const onArticleSubmit = useCallback(() => {
        if (inSubmit.current) return
        inSubmit.current = true
        // 检查输入
        if (form.noun.length === 0 || form.description.length === 0) {
            setErrorObj({msg: translation("article.header.errMsg")})
            inSubmit.current = false
            return
        }
        // 提交
        const createTime = dayjs().valueOf()
        setArticle(draft => {
            inSubmit.current = false
            hide()
            return [{
                createTime,
                title: form.noun,
                content: form.description,
                records: []
            }, ...draft.filter(item => item.createTime !== createTime)]
        })
    }, [form, setArticle, setErrorObj, hide])


    return <>
        <div
            font='medium'
            text='lg shadow-md fuchsia-800'
            m='y-1'
        >{translation(`${inquiryType}.header.add.title`)}
        </div>
        <ErrorTip obj={errorObj}/>
        <div m='y-2.5' space='y-2'>
            <div
            >
                <div
                    font='medium'
                    text='sm sky-700'
                    m='b-1'
                >{translation(`${inquiryType}.header.add.form`)[0].label}
                </div>
                <Input
                    value={form.noun}
                    onChange={set.noun!}
                    placeholder={translation(`${inquiryType}.header.add.form`)[0].placeholder}
                    text='placeholder-sky-700/30 placeholder-italic'
                />
            </div>
            <div
            >
                <div
                    font='medium'
                    text='sm sky-700'
                    m='b-1'
                >{translation(`${inquiryType}.header.add.form`)[1].label}
                </div>
                <Textarea
                    value={form.description}
                    onChange={set.description}
                    rows={window.innerWidth >= 768 ? inquiryType === 'article' ? 16 : 8 : inquiryType === 'article' ? 8 : 4}
                    placeholder={translation(`${inquiryType}.header.add.form`)[1].placeholder}
                    text='placeholder-sky-700/30 placeholder-italic'
                />
            </div>
        </div>
        <div
            flex='~ row-reverse'
            align='items-center'
            space='x-2.5 x-reverse'
        >
            <Button
                color='#a21caf'
                onClick={inquiryType === 'article' ? onArticleSubmit : onAdviceSubmit}
            >
                {translation('inquiry.header.add.submit')}
            </Button>
            <Button
                color='#a21caf'
                onClick={hide}
                variant='subtle'
            >
                {translation('inquiry.header.add.cancel')}
            </Button>
        </div>
    </>
}

const Header = (
    props: {
        loadedItem: LoadedType,
        setLoadedItem: Updater<LoadedType>
        inquiryType: InquiryType
    }
) => {

    const {
        loadedItem,
        setLoadedItem,
        inquiryType
    } = props

    const {translation} = useI18n()
    const buttonRef = useRef<HTMLButtonElement>(null)
    const {
        visible,
        hide,
        show
    } = useVisible(false)

    const iconRef = useRef(<Icon icon={contentSavePlusOutline}/>)

    return (
        <div flex='~' m='y-2' items='center'>
            <span
                select='none'
                font='medium'
                text='xl shadow-md fuchsia-900'
                flex='1'
            >{translation(`${inquiryType}.header.title`)}
            </span>
            <Button
                loading={loadedItem.id !== null}
                ref={buttonRef}
                color='#86198F'
                onClick={show}
                icon={iconRef.current}>
                {translation('inquiry.header.button')}
            </Button>
            <Dialog visible={visible} hide={hide} triggerRef={buttonRef}>
                <AddItemModal hide={hide} setLoadedItem={setLoadedItem} inquiryType={inquiryType}/>
            </Dialog>
        </div>
    )
}

export default Header
