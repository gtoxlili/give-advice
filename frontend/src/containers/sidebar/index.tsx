import "./style.css"
import React, {useCallback, useState} from "react";
import Tabs from "@components/tabs";

// @ts-ignore
import logo from "@assets/logo.png";

import {useI18n} from "@lib/hook";
import {useLocation, useNavigate} from "react-router-dom";


interface SidebarProps {
    routes: Array<{
        index?: boolean,
        path?: string,
    }>
}

export default function Sidebar(props: SidebarProps) {

    const {routes} = props
    const {pathname} = useLocation()
    const [tabIndex, setTabIndex] = useState(
        () => {
            const index = routes.findIndex(route => route.path === pathname)
            return index === -1 ? 0 : index
        })
    const {translation} = useI18n()
    const navigate = useNavigate()

    const changeTab = useCallback((index: number) => {
        setTabIndex(index)
        navigate(routes[index].path ?? '/')
    }, [setTabIndex])

    return (
        <div className="sidebar">
            <img src={logo} alt="logo" className="sidebar-logo"/>
            <span className='flex-1'/>
            <Tabs index={tabIndex}
                  tabsTitle={translation('sideBar.tabsTitle')}
                  setIndex={changeTab}
            />
        </div>
    )
}
