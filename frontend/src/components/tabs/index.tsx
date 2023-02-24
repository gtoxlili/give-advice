import React, {memo, useLayoutEffect} from "react";
import './style.css'
import classnames from "classnames";
import {useImmer} from "use-immer";

interface TabsProps {
    index: number
    tabsTitle: string[]
    setIndex: (index: number) => void
}

// 划块 Style
interface activeStyle {
    width: number
    offset: number
}

const TabActive = memo((props: { activeStyle: activeStyle }) => {
    const {activeStyle} = props

    return (
        <div
            pos='absolute bottom-0'
            bg='fuchsia-800'
            h='1'
            className='transition-active'
            style={{
                transform: `translateX(${activeStyle.offset}px)`,
                width: `${activeStyle.width}px`
            }}/>
    )
})

const TabContent = memo((props: TabsProps) => {
    const {index, tabsTitle, setIndex} = props

    return <div
        space='x-2'
    >
        {
            tabsTitle.map((title, i) => (
                <div
                    key={i}
                    id={`tabs-title-${i}`}
                    text='sm fuchsia-700'
                    p='x-3'
                    m='b-2'
                    display='inline-block'
                    select='none'
                    cursor='pointer'
                    className={classnames({
                        'font-bold': i === index,
                        'font-medium': i !== index,
                    },{'text-opacity-70': i !== index})}
                    onClick={() => {
                        setIndex(i)
                    }}>
                    {title}
                </div>
            ))
        }
    </div>
})

const Tabs = (props: TabsProps) => {
    const {index, tabsTitle} = props
    const [activeStyle, setActiveStyle] = useImmer<activeStyle>({} as activeStyle)

    useLayoutEffect(() => {
        const dom = document.getElementById(`tabs-title-${index}`)
        setActiveStyle(draft => {
            draft.offset = dom?.offsetLeft || 0
            draft.width = dom?.offsetWidth || 0
        })
    }, [index, tabsTitle])

    return (
        <div
            pos='relative'
        >
            <TabActive activeStyle={activeStyle}/>
            <TabContent {...props}/>
        </div>
    )
}

export default Tabs
