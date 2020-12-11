import {AxiosRequestConfig} from "axios";
import getInstance          from "@/api/instance";

const baseURL: string = "/";

// 请求配置
const axiosRequestConfig: AxiosRequestConfig = {
    baseURL,
    timeout: 30 * 1000,
};

// 请求接口
const instance = getInstance(axiosRequestConfig);

const Api = {

}

export default instance;
export {
    baseURL,
    instance,
    Api
}
