<template>
	<div ref="frame" style="height: 100vh">
		<iframe class="frame" src="http://192.168.3.2:9000"></iframe>
		<!--<iframe class="frame" src="http://192.168.1.100:9000"></iframe>-->
	</div>
</template>

<script lang="ts">
	import {defineComponent, ref, onMounted, getCurrentInstance} from "vue";
	import useNavBar                                             from "@/components/nav-bar/UseNavBar";

	export default defineComponent({
		name: "Portainer",
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

			onMounted(() => {
				ctx.$notify({
					title: '登录账号密码',
					message: '' +
						'用户名：admin<br>' +
						'密码：12345678',
					dangerouslyUseHTMLString: true,
					position: 'bottom-right',
					duration: 0
				});

			});

			return {
				frame
			};
		},
		methods: {}
	});
</script>

<style scoped>
	.frame {
		border: 0;
		width: 100%;
		height: 100%;
		overflow: hidden;
	}
</style>