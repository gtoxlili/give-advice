import {defineConfig, splitVendorChunkPlugin} from 'vite'
import react from '@vitejs/plugin-react'
import windiCSS from 'vite-plugin-windicss'
import tsConfigPath from 'vite-tsconfig-paths'
import jotaiDebugLabel from 'jotai/babel/plugin-debug-label'
import jotaiReactRefresh from 'jotai/babel/plugin-react-refresh'
import {VitePWA} from "vite-plugin-pwa";

export default defineConfig(env => {
    return {
        plugins: [// only use react-fresh
            env.mode === 'development' && react({
                babel: {plugins: [jotaiDebugLabel, jotaiReactRefresh]},
            }),
            windiCSS(),
            tsConfigPath(),
            splitVendorChunkPlugin(),
            VitePWA({
                registerType: 'autoUpdate', // 自动更新
                injectRegister: 'inline', // 注入到html中
                manifest: {
                    icons: [{
                        src: '//cdn.jsdelivr.net/gh/gtoxlili/give-advice/frontend/src/assets/logo.png',
                        sizes: '128x128',
                        type: 'image/png',
                    }],
                    start_url: '/',
                    short_name: 'Give Advice',
                    name: 'Give Advice',
                },
            })
        ],
        server: env.mode === 'development' ? {
            port: 3000,
            proxy: {
                '/api': {
                    target: 'http://localhost:7458',
                    changeOrigin: true,
                }
            }
        } : {},
        base: './',
        build: {
            reportCompressedSize: false,
        }
    }
})
