import {AxiosRequestConfig} from "axios";
import getInstance          from "@/api/instance";
import {Get}                from "@/api/http-method";

const baseURL: string = "/base";

// 请求配置
const axiosRequestConfig: AxiosRequestConfig = {
    baseURL,
    timeout: 30 * 1000,
};

// 请求接口
const instance = getInstance(axiosRequestConfig);

const api = {
    apps:"/apps/"
}

const baseApi = {
    getApps(){
        return Get(instance,api.apps);
    }
}


export default baseApi;
export {
    baseURL,
    instance,
    baseApi
}
