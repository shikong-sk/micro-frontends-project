import {AxiosRequestConfig} from "axios";
import getInstance          from "@/api/instance";

// 请求配置
const axiosRequestConfig: AxiosRequestConfig = {
    baseURL: "/",
    timeout: 30 * 1000,
};

// 请求接口
const instance = getInstance(axiosRequestConfig);

export default instance;
