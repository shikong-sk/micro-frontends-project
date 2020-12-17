import {createStore, ModuleTree} from 'vuex';

let modules: ModuleTree<{}> = {};
let modulesFiles = require.context("./modules", true, /[\.js,\.ts]$/);
console.log("modulesFiles", modulesFiles, modulesFiles.keys());
modulesFiles.keys().reduce((modules, modulePath) => {
	// set './app.js' => 'app'
	const moduleName: string = modulePath.replace(/^\.\/(.*)\.\w+$/, '$1');
	const value = modulesFiles(modulePath);
	modules[moduleName] = value.default;
	return modules;
}, modules);
console.log(modules);

export default createStore({
	modules,
	state:{},
	mutations: {},
	actions: {},
});
