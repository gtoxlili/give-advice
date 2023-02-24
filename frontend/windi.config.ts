import {defineConfig} from 'vite-plugin-windicss'

export default defineConfig({
    attributify: true,
    theme: {
        extend: {}
    },
    plugins: [
        require('windicss/plugin/line-clamp')
    ]
})
