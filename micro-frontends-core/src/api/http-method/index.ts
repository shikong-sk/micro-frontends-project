import {AxiosInstance, AxiosRequestConfig} from "axios";

const httpMethod = {
    // get请求
    Get(request: AxiosInstance, url: string, params: {} = {}, config: AxiosRequestConfig = {}) {
        config = config || {};
        config.params = params;
        return request.get(url, config);
    },
    // delete请求
    Delete(request: AxiosInstance, url: string, params: {} = {}, config: AxiosRequestConfig = {}) {
        config.params = params;
        return request.delete(url, config);
    },
    // post请求
    Post(request: AxiosInstance, url: string, params: {} = {}, config: AxiosRequestConfig = {}) {
        return request.post(url, params, config);
    },
    // put请求
    Put(request: AxiosInstance, url: string, params: {} = {}, config: AxiosRequestConfig = {}) {
        return request.put(url, params, config);
    },
    // patch请求
    Patch(request: AxiosInstance, url: string, params: {} = {}, config: AxiosRequestConfig = {}) {
        return request.patch(url, params, config);
    }
};

const {Get, Post, Put, Delete, Patch} = httpMethod;
export {Get, Post, Put, Delete, Patch};
export default httpMethod;
