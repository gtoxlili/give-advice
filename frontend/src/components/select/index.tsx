import classnames from 'classnames'
import './style.css'
import React from "react";

export interface SelectOptions<T> {
    label: string
    value: T
}

export interface SelectProps<T> {
    // options
    options: Array<SelectOptions<T>>

    // active value
    value: T

    // select callback
    onSelect: (value: T) => void
}

export function Select<T>(props: SelectProps<T>) {
    const {options, value, onSelect} = props

    return (
        <div
            flex='~ row'
            font='medium'
            text='xs'
            className="divide-x-2">
            {
                options.map(option => (
                    <button
                        key={option.label}
                        className={classnames(
                            'px-3', 'button-select-options',
                            'last:pr-0', 'first:pl-0',
                            {'text-slate-400': value !== option.value,},
                            {'text-sky-700 font-semibold text-sm': value === option.value}
                        )}
                        onClick={() => onSelect?.(option.value)}>
                        {option.label}
                    </button>
                ))
            }
        </div>
    )
}
