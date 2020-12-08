import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

createApp(App).use(store).use(router).mount('#app')

// 乾坤微前端框架
import { registerMicroApps, start }     from 'qiankun';
const apps = [
	{
		name:"vue",
		entry: "http://localhost:10001/",
		container:"#vue",
		activeRule:"/vue"
	}
]

registerMicroApps(apps);

start();
