//  使用人次
import useSWR from "swr";
import {useClient} from "@lib/hook";

export const useUserCount = () => {
    const client = useClient()
    const {data, error} = useSWR(
        ["info/useCount", client], async ([_, client]) => {
            const res = await client.getUsageCount()
            return res.data.count
        }, {keepPreviousData: true}
    )
    return [data, error]
}



