import React, {ReactNode} from "react";
import Footer from "@containers/inquiry/components/footer";
import {LanguageSelect, OwnToken, ProjectLink} from "@containers/settings/component";
import {useI18n} from "@lib/hook";

const Header = (
    {title}: { title: string }
) => {
    return <div
        m='y-3'
        select='none'
        font='medium'
        text='xl shadow-md fuchsia-900'
    >{title}
    </div>
}

interface SettingModelProps {
    title: string
    children: ReactNode
}

const SettingModel = (props: SettingModelProps) => {
    const {title, children} = props
    return <div>
        <div
            text='color-[#54759a] sm'
            font='semibold'
            m='b-1 l-1'
        >
            {title}
        </div>
        <div
            shadow='~'
            border='rounded'
            p='y-4 x-8'
            bg='white'
            space='y-4'
        >
            {children}
        </div>
    </div>
}

const Settings = () => {
    const {translation} = useI18n()

    return <div flex='~ col' h='full'>
        <Header title={translation("settings.title")}/>
        <div flex='1 ~ col' overflow='y-auto' p='x-3' m='-x-3'>
            <div flex='1' m='y-2.5' space='y-4'>
                <SettingModel title={translation("settings.first.title")}>
                    <LanguageSelect/>
                    <OwnToken/>
                </SettingModel>
                <SettingModel title={translation("settings.second.title")}>
                    <ProjectLink/>
                </SettingModel>
            </div>
            <Footer/>
        </div>
    </div>
}

export default Settings
