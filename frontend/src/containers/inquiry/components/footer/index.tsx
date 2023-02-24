import React from 'react'
import "./style.css"
import {useI18n} from "@lib/hook";
import {useUserCount} from "@stores/swr";

const Count = () => {
    const {translation} = useI18n()
    const [count, err] = useUserCount()
    return err ? <></> : <>
        <div>
            {translation("inquiry.footer.two.left")}
            <span className="mx-1 font-medium text-gray-500">{count}</span>
            {translation("inquiry.footer.two.right")}
        </div>
        <div className="mx-2">|</div>
    </>
}

const Footer = () => {
    const {translation} = useI18n()

    return <div
        text="xs gray-400"
        m='2'
        flex='~ row'
        align='items-center'
        justify='center'
        select='none'
    >
        <div>
            {translation("inquiry.footer.one.left")}
            <a
                className="font-medium text-gray-500 hover:text-gray-600 cursor-pointer mx-1"
                href="https://openai.com/"
                target="_blank"
            >OpenAI</a>
            {translation("inquiry.footer.one.right")}
        </div>
        <div className="mx-2">|</div>
        <Count/>
        <div>
            <a className="font-medium text-gray-500 hover:text-gray-600 cursor-pointer"
               href="mailto:gtoxlili@outlook.com">
                {translation("inquiry.footer.three")}
            </a>
        </div>
    </div>
}
export default Footer
