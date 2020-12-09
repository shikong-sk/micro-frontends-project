import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from "element-plus";
import 'element-plus/lib/theme-chalk/index.css';

const app = createApp(App).use(store).use(router)
app.use(ElementPlus)
app.mount('#app')

// 乾坤微前端框架
import { registerMicroApps, start,loadMicroApp }     from 'qiankun';
const apps = [
	{
		name:"vue",
		entry: "http://192.168.1.100:10001/",
		container:"#vue",
		activeRule:"/vue"
	},{
		name:"ncda",
		entry: "http://192.168.1.100:10002/ncda/index",
		container:"#ncda",
		activeRule:"/micro-ncda"
	}
]


const excludeAssets:Array<string> = [
	"https://api.map.baidu.com/getscript?v=3.0&ak=DCfcc8d7f591e4106fb6d5d7390df429",
	"http://api.map.baidu.com/library/GeoUtils/1.2/src/GeoUtils_min.js",
	"http://api.map.baidu.com/library/Heatmap/2.0/src/Heatmap_min.js"
];

const customImportConfig = {
	excludeAssetFilter(url:string){
		for(let i=0;i < excludeAssets.length;i++){
			let excludeAssetsURL = excludeAssets[i];
			if(url === excludeAssetsURL){
				return true;
			}
		}
		return false;
	},
	// @ts-ignore
	fetch(url, ...args) {
		for(let i=0;i < excludeAssets.length;i++){
			let excludeAssetsURL = excludeAssets[i];
			if(url === excludeAssetsURL){
				insertEle(url);
				return Promise.resolve({
					status: 200,
					text:()=>{return ""}
				});
			}
		}
		return window.fetch(url, ...args);
	}
};

// @ts-ignore
function insertEle(url){
	let head = document.getElementsByTagName('head')[0]
	let el = document.createElement('script')
	let id = +new Date() + Math.random()

	el.src = url
	el.id = String(id)
	head.append(el)
	el.onload = function(){
		let exist = false;
		for(let i=0;i < excludeAssets.length;i++){
			let excludeAssetsURL = excludeAssets[i];
			if(url === excludeAssetsURL){
				exist = true
			}
		}
		if(!exist){
			// @ts-ignore
			head.removeChild(document.getElementById(id))
		}
	}
}

// @ts-ignore
registerMicroApps(apps,	{});

// @ts-ignore
start(customImportConfig);
