<template>
	<div id="nav">
		<router-link to="/">Home</router-link> |
		<router-link to="/about">About</router-link> |
		<router-link to="/vue">子应用</router-link>
	</div>
	<div class="about">
		<h1>This is an about page</h1>
	</div>
	<div ref="sub"></div>
</template>

<script lang="ts">
	import { loadMicroApp } from "qiankun";
	import {defineComponent} from "vue";

	export default defineComponent({
		mounted() {
			this.$nextTick(() => {
				// 在内部页面作为某个组件动态引入
				let app = loadMicroApp({
					name: "subApp",
					entry: "http://localhost:10001/vue/about",
					// @ts-ignore
					container: this.$refs.sub
				});

				console.log(app);
			});
		}
	})
</script>
