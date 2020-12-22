<template>
	<div ref="frame" style="height: 100vh">
		<iframe class="frame" :src="target"></iframe>
	</div>
</template>

<script lang="ts">
	import {defineComponent, ref, onMounted, getCurrentInstance} from "vue";
	import useNavBar                                             from "@/components/nav-bar/UseNavBar";
	import {RouteRecord}                                         from "vue-router";

	export default defineComponent({
		name: "MinIO",
		setup: (prop, context) => {
			let frame = ref(null);
			console.log(frame);
			useNavBar(frame);

			// 获取当前内部组件实例
			const internalInstance = getCurrentInstance();
			let ctx: Record<string, any> = {};
			if (internalInstance != null) {
				// 获取全局属性/方法
				ctx = internalInstance.appContext.config.globalProperties;
			}

			let target = ref<string>("http://192.168.3.2:20001 ");

			onMounted(() => {
				ctx.$notify({
					title: '认证参数',
					customClass: "notify-medium",
					message: '' +
						'Access Key: AKIAIOSFODNN7EXAMPLE<br>' +
						'Secret Key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY',
					dangerouslyUseHTMLString: true,
					position: 'bottom-right',
					duration: 0
				});
				let route = ctx.$route;
				let query = route.query;
				if (query.target) {
					target.value = "http://" + query.target;
				}
			});

			return {
				frame,
				target
			};
		},
	});
</script>

<style>
	.notify-medium {
		width: 500px;
	}
</style>
<style scoped>

	.frame {
		border: 0;
		width: 100%;
		height: 100%;
		overflow: hidden;
	}
</style>