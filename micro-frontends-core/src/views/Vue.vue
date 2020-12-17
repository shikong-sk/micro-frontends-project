<template>
	<div id="vue"></div>
</template>

<script>
	import {start} from "qiankun";
	import { globalStateActions } from "../main";

	export default {
		name: "Vue",
		mounted() {
			this.$nextTick(()=>{
				if (!window.qiankunStarted) {
					window.qiankunStarted = true;
					start();
				}
				globalStateActions.setGlobalState({"containerReady":true});
			})
		},
		unmounted() {
			globalStateActions.setGlobalState({"containerReady":false});
		}
	}
</script>

<style scoped>

</style>
