import {createStore} from 'vuex';

let modulesFiles = require.context("./modules", true, /[\.js,\.ts]$/);
console.log("modulesFiles", modulesFiles);

export default createStore({
	modules: {},
	state: {},
	mutations: {},
	actions: {},
});
