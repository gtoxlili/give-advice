import {createPortal} from "react-dom";
import {CSSProperties, memo, ReactNode, RefObject, useLayoutEffect, useMemo, useState} from "react";
import {useImmer} from "use-immer";
import './style.css'

interface MaskProps {
    onClick: () => void;
}

const Mask = memo((props: MaskProps) => {
    const {
        onClick
    } = props;
    return createPortal(<div
            pos='fixed top-0 left-0 right-0 bottom-0'
            z='3999'
            onClick={onClick}
        />
        , document.body)
})

interface IDialogProps {
    visible: boolean
    children: ReactNode
    position: CSSProperties
}

const IDialog = (props: IDialogProps) => {
    const {
        visible,
        children,
        position
    } = props

    const dialogStyle = useMemo(() => {
        return visible ? {
            opacity: 1,
            transform: window.innerWidth >= 768 ? 'scale(1)' : 'scale(1) translateY(-50%)',
        } satisfies CSSProperties : {
            opacity: 0,
            transform: 'scale(0)',
            ...window.innerWidth >= 768 ? {} : position
        } satisfies CSSProperties
    }, [visible, position])

    const [vis, setVis] = useState(visible)
    useLayoutEffect(() => {
        visible && setVis(true)
    }, [visible])

    return createPortal(
        <div
            z='4000'
            className='dialog'
            select='none'
            style={dialogStyle}
            onTransitionEnd={(e) => e.propertyName === 'opacity' && !visible && setVis(false)}
        >{vis && children}</div>, document.body
    )
}

interface DialogProps {
    visible: boolean
    hide: () => void
    children: ReactNode
    triggerRef: RefObject<HTMLElement>
}

const Dialog = (props: DialogProps) => {
    const {
        visible,
        hide,
        children,
        triggerRef
    } = props

    // dialogPosition
    const [dialogPosition, setDialogPosition] = useImmer<CSSProperties>({} as CSSProperties)
    useLayoutEffect(() => {
        if (!triggerRef.current) return
        const dom = triggerRef.current
        // 以 triggerRef 为中心点
        setDialogPosition(draft => {
            draft.top = dom.offsetTop + dom.offsetHeight / 2
            draft.right = window.innerWidth - dom.offsetLeft - dom.offsetWidth / 2
        })
    }, [])


    return <>
        {visible && <Mask onClick={() => window.innerWidth >= 768 && hide()}/>}
        <IDialog visible={visible} position={dialogPosition}>{children}</IDialog>
    </>
}

export default Dialog
