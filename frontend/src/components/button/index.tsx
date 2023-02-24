import React, {CSSProperties, forwardRef, memo, ReactNode, Ref, useMemo} from "react";
import classnames from "classnames";
import './style.css'
import {colorOpacity, noop} from "@lib/helper";

interface ButtonProps {
    disabled?: boolean
    fullWidth?: boolean
    loading?: boolean
    color: string
    children?: ReactNode
    onClick?: () => void
    icon?: ReactNode
    uppercase?: boolean
    variant?: 'filled' | 'outlined' | 'light' | 'subtle'
}

const Button = (props: ButtonProps, ref: Ref<HTMLButtonElement>) => {
    const {
        disabled,
        fullWidth,
        loading,
        color,
        children,
        onClick = noop,
        icon,
        uppercase,
        variant = 'filled'
    } = props

    const btnStyle = useMemo(() => {
        switch (variant) {
            case 'filled':
                return {
                    backgroundColor: color,
                    color: 'white',
                    borderColor: 'transparent',
                    '--hover-bg-color': color,
                } satisfies CSSProperties
            case 'outlined':
                return {
                    borderColor: color,
                    color: color,
                    backgroundColor: 'transparent',
                    '--hover-bg-color': colorOpacity(color, 0.05),
                } satisfies CSSProperties
            case 'light':
                return {
                    backgroundColor: colorOpacity(color, 0.1),
                    color: color,
                    borderColor: 'transparent',
                    '--hover-bg-color': colorOpacity(color, 0.2),
                } satisfies CSSProperties
            default:
                return {
                    backgroundColor: 'transparent',
                    color: color,
                    borderColor: 'transparent',
                    '--hover-bg-color': colorOpacity(color, 0.2),
                } satisfies CSSProperties
        }
    }, [color, variant])


    return <button
        disabled={disabled || loading}
        className={
            classnames(
                'btn',
                'text-sm',
                'font-medium',
                'tracking-widest',
                {'w-full': fullWidth},
                {'uppercase': uppercase},
                icon ? 'pr-4' : 'px-4',
            )}
        style={btnStyle}
        onClick={onClick}
        ref={ref}
    >
        {
            loading ? <
                svg
                className="animate-spin" xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
            >
                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"/>
                <path className="opacity-75" fill="currentColor"
                      d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
            </svg> : icon
        }
        {children}
    </button>
}

export default memo(forwardRef(Button))
