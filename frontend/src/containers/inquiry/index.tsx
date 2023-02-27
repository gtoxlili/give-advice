import React, {useMemo} from 'react'

import Header from "@containers/inquiry/components/header";

import {AdviceType as A} from "@stores/jotai";
import {useImmer} from "use-immer";
import Footer from "@containers/footer";
import {useLocation} from "react-router-dom";
import {InquiryType} from "@containers/app";
import Advice from "@containers/inquiry/components/advice";
import Article from "@containers/inquiry/components/article";

export type LoadedType = Omit<A, "answer"> & { id: string | null, error: boolean }

const Inquiry = () => {

    const [loadedItem, setLoadedItem] = useImmer<LoadedType>({id: null, error: false} as LoadedType)
    // 获取当前路由
    const {pathname} = useLocation()
    const inquiryType = useMemo(() => {
        return pathname === '/' ? 'advice' : pathname.slice(1) as InquiryType
    }, [pathname])

    return <div flex='~ col' h='full'>
        <Header loadedItem={loadedItem} setLoadedItem={setLoadedItem} inquiryType={inquiryType}/>
        <div flex='1 ~ col' overflow='y-auto' p='x-3' m='-x-3'>
            {inquiryType === 'advice' && <Advice loadedItem={loadedItem} setLoadedItem={setLoadedItem}/>}
            {inquiryType === 'article' && <Article/>}
            <Footer/>
        </div>
    </div>
}

export default Inquiry
