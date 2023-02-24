import React, {StrictMode} from 'react'

import 'virtual:windi.css'
import '@styles/common.css'
import {createRoot} from "react-dom/client";
import App from "@containers/app";


const AppInstance = (
    <StrictMode>
        <App/>
    </StrictMode>
)

createRoot(document.getElementById('root')!).render(AppInstance)
