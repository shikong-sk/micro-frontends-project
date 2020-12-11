import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse} from 'axios';

// 创建请求 并返回一个请求实例
export function getInstance(axiosRequestConfig: AxiosRequestConfig) : AxiosInstance {
    let instance = axios.create(axiosRequestConfig);
    // 全局 设置拦截器
    instance.interceptors.response.use(successCB, rejectCB);
    return instance;
}

// 请求成功回调
function successCB(response: AxiosResponse | Promise<AxiosResponse>): AxiosResponse | Promise<AxiosResponse> {
    function processResponse(response: AxiosResponse) {
        if (response.status === 200) { // 响应成功
            return response.data;
        } else {
            return Promise.reject(response);
        }
    }

    if (response instanceof Promise) {
        let resp: AxiosResponse = {
            config    : {},
            data      : undefined,
            headers   : undefined,
            status    : 500,
            statusText: ""
        };
        response.then((r: AxiosResponse) => {
            resp = r;
        });
        return processResponse(resp);
    } else {
        return processResponse(response);
    }
}

// 请求失败回调
function rejectCB(error: any): any {
    return Promise.reject(error);
}

export default getInstance;
