import React, {memo, TextareaHTMLAttributes} from "react";
import './style.css'
import {noop} from "@lib/helper";
import classnames from "classnames";

type TextareaProps =
    Omit<TextareaHTMLAttributes<HTMLTextAreaElement>, 'onChange'>
    & { onChange?: (value: string) => void }

const Textarea =
    (props: TextareaProps) => {
        const {
            value = '',
            rows,
            className,
            onChange = noop,
            ...attributes
        } = props

        return (
            <textarea
                value={value}
                onChange={event => onChange(event.target.value)}
                rows={rows ? rows : window.innerWidth < 768 ? 4 : 16}
                className={classnames(className, 'textarea-text focus:shadow-md focus:shadow-sky-800/10')}
                {...attributes}
            />
        )
    }

export default memo(Textarea)
