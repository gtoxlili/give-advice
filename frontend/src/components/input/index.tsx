import React, {FocusEvent, forwardRef, KeyboardEvent, memo, ReactNode, Ref, useCallback} from "react";
import './style.css'
import {noop} from "@lib/helper";
import {AttributifyAttributes} from "windicss/types/jsx";

interface InputProps extends AttributifyAttributes {
    autoFocus?: boolean
    disabled?: boolean
    onBlur?: (event?: FocusEvent<HTMLInputElement>) => void
    onChange: (value: string) => void
    onEnter?: (event?: KeyboardEvent<HTMLInputElement>) => void
    onFocus?: (event?: FocusEvent<HTMLInputElement>) => void
    placeholder?: string
    suffixIcon?: ReactNode
    value?: string
}

const Input = forwardRef(
    (props: InputProps, ref: Ref<HTMLDivElement>) => {

        const {
            value = '',
            onChange,
            onEnter = noop,
            suffixIcon,
            ...attributes
        } = props

        const handleKeyDown = useCallback(
            (e: KeyboardEvent<HTMLInputElement>) => {
                if (e.code === 'Enter') {
                    onEnter(e)
                }
            }, [onEnter])

        return (
            <div
                ref={ref}
                pos='relative'
            >
                <div className='suffixIcon'>
                    {suffixIcon}
                </div>
                <input
                    value={value}
                    onChange={event => onChange(event.target.value)}
                    onKeyDown={handleKeyDown}
                    className='input-text focus:shadow-md focus:shadow-sky-800/10'
                    {...attributes}
                />
            </div>
        )
    })

export default memo(Input)
