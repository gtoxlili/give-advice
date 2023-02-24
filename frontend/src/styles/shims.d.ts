import {AttributifyAttributes} from "windicss/types/jsx";

declare module "react" {

    interface HTMLAttributes<T> extends AttributifyAttributes {
    }

    interface CSSProperties {
        // 自定义属性
        '--hover-bg-color'?: string
    }
}
