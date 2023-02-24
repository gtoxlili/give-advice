import React, {lazy, Suspense} from "react";
import Sidebar from "@containers/sidebar";
import {createBrowserRouter, createHashRouter, Outlet, RouterProvider} from "react-router-dom";
import "./style.css";

const Inquiry = lazy(() => import('@containers/inquiry'))
const Settings = lazy(() => import('@containers/settings'))

export type InquiryType = 'advice' | 'article'

const routers = [
    {
        path: "/",
        element: <Index/>,
        children: [
            {
                index: true,
                element: <Inquiry/>,
            },
            {
                path: "/article",
                element: <Inquiry/>,
            },
            {
                path: "/setting",
                element: <Settings/>,
            },
        ]
    },
    {
        path: "/login",
        element: <p>coming soon</p>,
    }
]

const routerInstance = createHashRouter(routers)

function Index() {
    return <div className="app">
        <Sidebar routes={routers[0].children!}/>
        <div className="app-container">
            <Outlet/>
        </div>
    </div>
}

const App = () => <Suspense><RouterProvider router={routerInstance}/></Suspense>

export default App
