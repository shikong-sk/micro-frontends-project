<template>
	<div id="vue" ref="container"></div>
</template>

<script lang="ts">
	import {defineComponent, ref, onMounted, onUnmounted, inject, watch} from "vue";
	import {start}                                                       from "qiankun";
	import {globalStateActions}                                          from "@/main";
	import useNavBar                                from "@/components/nav-bar/useNavBar";

	export default defineComponent({
		name: "Vue",
		setup(props, context) {
			let collapseNav = ref(inject("collapseNav"));
			watch(collapseNav, (val) => {
				console.log("collapseNav - vue", val);
			});

			let container = ref();
			useNavBar(container, 20);


			onMounted(() => {
				// @ts-ignore
				if (!window.qiankunStarted) {
					// @ts-ignore
					window.qiankunStarted = true;
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