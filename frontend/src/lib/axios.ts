import axios, {AxiosError, AxiosInstance} from 'axios'
import {InquiryType} from "@containers/app";

export interface RegisterTopic {
    noun: string
    description: string
}

export interface Res<T> {
    code: number
    message: string
    data: T
}


export class Client {
    private readonly axiosClient: AxiosInstance

    constructor(url: string, secret?: string, token?: string) {
        const headers = {
            Authorization: `Bearer ${secret}`,
            "OpenAI-Auth-Key": `${token}`
        }
        this.axiosClient = axios.create({
            baseURL: url,
            headers: headers,
        })
    }

    // 注册一个主题
    async registerTopic(type: InquiryType, topic: RegisterTopic) {
        try {
            const res = await this.axiosClient.post<Res<{ id: string }>>('register/' + type, topic)
            return res.data
        } catch (e) {
            const res = e as AxiosError
            console.log(res)
            return {
                code: res.status,
                message: res.response ? res.response.statusText : res.message,
                data: null
            }
        }
    }

    // 获取总使用人数
    async getUsageCount() {
        const res = await this.axiosClient.get<Res<{ count: number }>>('info/useCount')
        return res.data
    }

}
