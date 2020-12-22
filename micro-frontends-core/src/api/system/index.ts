import {AxiosRequestConfig} from "axios";
import getInstance          from "@/api/instance";
import {Get}                from "@/api/http-method";

const baseURL: string = "/system";

// 请求配置
const axiosRequestConfig: AxiosRequestConfig = {
	baseURL,
	timeout: 30 * 1000,
};

// 请求接口
const instance = getInstance(axiosRequestConfig);

const api = {
	getNetWorkApi:"/getNetWork"
}

const systemApi = {
	getNetWork(){
		return Get(instance,api.getNetWorkApi);
	}
}

export default systemApi;
export {
	baseURL,
	instance,
	systemApi
}
