import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 原理是把这个文件打包成一个类库，让微前端的基座引用
let instance: { $destroy: () => void } | null = null
function render(props = {}){
	// @ts-ignore
	const {container} = props;
	//  这个传入的props 可以通讯使用
	// @ts-ignore
	instance = createApp(App).use(store).use(router).mount(container?container.querySelector('#app'):'#app');
}
//  动态设置 文件路径
// @ts-ignore
if(window.__POWERED_BY_QIANKUN__){
	// @ts-ignore
	console.log(window.__POWERED_BY_QIANKUN__);
	// @ts-ignore
	__webpack_public_path__ = window.__INJECTED_PUBLIC_PATH_BY_QIANKUN__;
}
//  独立运行项目
// @ts-ignore
if(!window.__POWERED_BY_QIANKUN__){
	render()
}
//  这个方式内容可以不写 ，但是必须定义 内部会校验 必须返回一个promise
// @ts-ignore
export async function bootstrap(props){
	//  启动
	console.log('bootstrap',props);
}

//  父应用传过来的api
// @ts-ignore
export async function mount(props){
	render(props)
	console.log('mount',props);
}

//  组件卸载时候
export async function unmount(){
	//  卸载组件
	// @ts-ignore
	instance.$destroy()
}
