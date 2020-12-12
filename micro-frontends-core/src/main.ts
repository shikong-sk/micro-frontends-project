import {createApp} from 'vue';
import App         from './App.vue';
import router      from './router';
import store       from './store';
import ElementPlus from "element-plus";
import 'element-plus/lib/theme-chalk/index.css';

import api from "@/api";

const app = createApp(App).use(store).use(router);
// 全局注册 API
app.config.globalProperties.$api = api;
app.use(ElementPlus);
app.mount('#app');
console.log(app);

// 乾坤微前端框架
import {
	FrameworkConfiguration,
	FrameworkLifeCycles,
	LoadableApp,
	registerMicroApps,
	RegistrableApp,
	start
}          from 'qiankun';

// 排除的 特殊的动态加载的微应用资源 JavaScript 资源 , 不被 qiankun 劫持处理
declare type ExcludeJavaScriptAsset = { url: string, async?: boolean, defer?: boolean };
// 排除的 特殊的动态加载的微应用资源 JavaScript 资源 列表
declare type ExcludeJavaScriptAssetList = Array<ExcludeJavaScriptAsset>;

// 排除的 特殊的动态加载的微应用资源 JavaScript 资源 列表
const excludeJavaScriptAssets: ExcludeJavaScriptAssetList = [
	{
		url: "https://api.map.baidu.com/getscript?v=3.0&ak=DCfcc8d7f591e4106fb6d5d7390df429"
	},
	{
		url: "http://api.map.baidu.com/library/GeoUtils/1.2/src/GeoUtils_min.js",
		defer: true
	},
	{
		url: "http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js",
		defer: true
	},
	{
		url: "http://api.map.baidu.com/?qt=verify&ak=DCfcc8d7f591e4106fb6d5d7390df429&callback=BMap._rd._cbk74925"
	},
	{
		url: "http://api.map.baidu.com/?qt=business_accredit&ak=DCfcc8d7f591e4106fb6d5d7390df429&callback=BMap._rd._cbk94067"
	}
];

const customImportConfig: FrameworkConfiguration = {
	/**
	 * 沙箱 配置
	 * 默认为 true
	 * <p/>
	 * 可设为 strictStyleIsolation 和 experimentalStyleIsolation
	 * <p/>
	 * strictStyleIsolation 开启严格的样式隔离模式
	 * 这种模式下 qiankun 会为每个微应用的容器包裹上一个 shadow dom 节点，从而确保微应用的样式不会对全局造成影响
	 * 已知对部分生成的组件的支持有问题
	 * <p/>
	 * experimentalStyleIsolation 会改写子应用所添加的样式为所有样式规则增加一个特殊的选择器规则来限定其影响范围
	 * 如 div .app-main => div[app应用名称] .app-main
	 * 已知对部分生成的组件的支持有问题
	 */
	sandbox: true,
	// 指定部分特殊的动态加载的微应用资源（css/js) 不被 qiankun 劫持处理
	excludeAssetFilter(url: string) {
		for (let i = 0; i < excludeJavaScriptAssets.length; i++) {
			let excludeJavaScriptAsset = excludeJavaScriptAssets[i];
			if (url === excludeJavaScriptAsset.url) {
				return true;
			}
		}
		return false;
	},
	// 自定义 qiankun 的 fetch 方法
	fetch(input: RequestInfo, init?: RequestInit | undefined) {
		// 拦截 排除的 js 资源加载
		for (let i = 0; i < excludeJavaScriptAssets.length; i++) {
			let excludeJavaScriptAsset = excludeJavaScriptAssets[i];
			if (input === excludeJavaScriptAsset.url) {
				insertScriptElement(input);
				return Promise.resolve(new Response());
			}
		}
		return window.fetch(input, init);
	}
};

// 插入 拦截的 Script 脚本 到主框架的 header 中
function insertScriptElement(url: RequestInfo) {
	let head = document.getElementsByTagName('head')[0];
	let el = document.createElement('script');
	let id = +new Date() + Math.random();

	if (typeof url === "string") {
		for (let i = 0; i < excludeJavaScriptAssets.length; i++) {
			let excludeJavaScriptAsset = excludeJavaScriptAssets[i];
			if (url === excludeJavaScriptAsset.url) {
				el.src = url;
				let {async,defer} = excludeJavaScriptAsset;
				el.async = async || false;
				el.defer = defer || false;
				break;
			}
		}
	}

	el.id = String(id);
	head.append(el);
	el.onload = function () {
		let exist = false;
		for (let i = 0; i < excludeJavaScriptAssets.length; i++) {
			let excludeJavaScriptAsset = excludeJavaScriptAssets[i];
			if (url === excludeJavaScriptAsset.url) {
				exist = true;
				break;
			}
		}
		if (!exist) {
			let targetNode: HTMLElement | null = document.getElementById(String(id));
			if(targetNode){
				head.removeChild(targetNode);
			}
		}
	};
}

// 全局微应用 生命周期 钩子
const appLifeCycles: FrameworkLifeCycles<any> = {
	// 准备加载
	beforeLoad: (): Promise<any> => {
		return Promise.resolve();
	},

	// 准备挂载
	beforeMount: (): Promise<any> => {
		return Promise.resolve();
	},

	// 挂载完成
	afterMount: (app: LoadableApp<any>, global: Window): Promise<any> => {
		console.log("afterMount", app, global);
		return Promise.resolve();
	},

	// 准备卸载
	beforeUnmount: (): Promise<any> => {
		return Promise.resolve();
	},

	// 卸载完成
	afterUnmount: (): Promise<any> => {
		return Promise.resolve();
	}
};

let registerApps: RegistrableApp<any>[] = [];

new Promise<void>((resolve)=>{
	// 从 外部接口 载入 子应用资源信息
	api.baseApi.getConfig().then((response)=>{
		console.log(response);
		let res = response.data;
		let apps = res.apps || [];
		apps.forEach((app: RegistrableApp<any>)=>{
			registerApps.push(app)
		})
		resolve();
	})
}).then(()=>{
	registerMicroApps(registerApps, appLifeCycles);
	start(customImportConfig);
})

