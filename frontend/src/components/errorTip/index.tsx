import "./style.css"
import React, {memo, useEffect, useMemo, useRef} from "react";

interface ErrorTipProps {
    obj: { msg: string }
}

const ErrorTip = memo((props: ErrorTipProps) => {
    const {obj} = props

    // 用于存储上次的动画名
    const lastAnimationName = useRef<string>()
    useEffect(() => {
        lastAnimationName.current = lastAnimationName.current === 'shakeB' ? 'shakeA' : 'shakeB'
    }, [obj])

    const errorMsg = useMemo(() => {
        if (obj.msg) {
            return {
                msg: obj.msg,
                display: 'block',
                animationName: lastAnimationName.current,
            }
        } else {
            return {msg: '', display: 'none'}
        }
    }, [obj])

    return <p className="error-tip"
              font='medium'
              text='xs rose-700'
              style={errorMsg}
    >* {errorMsg.msg}</p>
})

export default ErrorTip
