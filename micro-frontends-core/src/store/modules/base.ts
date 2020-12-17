import {RegistrableApp} from "qiankun";

let registerApps: RegistrableApp<any>[] = [];
let state = {
	registerApps
};

let getters = {
	getRegistrableApps(state:any){
		return state.registerApps;
	}
};

let mutations = {
	setRegisterApps(state:any,args:any){
		console.log(state,args);
		state.registerApps = args;
	}
};

let actions = {};


export default {
	namespaced: true,
	state,
	getters,
	mutations,
	actions
};