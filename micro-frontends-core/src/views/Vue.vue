<template>
	<div id="vue" ref="container"></div>
</template>

<script lang="ts">
	import {defineComponent, ref, onMounted, onUnmounted, inject, watch} from "vue";
	import {start}                                                       from "qiankun";
	import {globalStateActions}                                          from "@/main";
	import UseNavBar                                                     from "@/components/nav-bar/UseNavBar";

	export default defineComponent({
		name: "Vue",
		setup(props, context) {
			let container = ref();
			UseNavBar(container, 20);

			let collapseNav = ref(inject("collapseNav"));
			watch(collapseNav, (val) => {
				console.log("collapseNav - vue", val);
			});

			onMounted(() => {

				if (!(window as any).qiankunStarted) {
					(window as any).qiankunStarted = true;
					start();
				}
				globalStateActions.setGlobalState({"containerReady": true});
			});
			onUnmounted(() => {
				globalStateActions.setGlobalState({"containerReady": false});
			});
			return {
				container
			};
		},
	});
</script>