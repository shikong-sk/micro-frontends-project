<template>
	<div class="nav-bar">
		<el-menu default-active="1-4-1" class="el-menu-vertical" @select="handleSelect" @open="handleOpen"
		         @close="handleClose" :collapse="isCollapse">
			<div>
				<el-button style="border: none;width: 100%"
				           :icon="isCollapse?'el-icon-s-unfold':'el-icon-s-fold'"
				           @click="collapseChange">
					{{isCollapse?'':'收起导航栏'}}
				</el-button>
			</div>
			<el-menu-item index="home">
				<i class="el-icon-s-home"></i>
				<template #title>主框架首页</template>
			</el-menu-item>
			<el-submenu index="Project">
				<template #title>
					<i class="el-icon-menu"></i>
					<span>微前端功能页</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="ncda">
						<template #title>NCDA 1.8 Beta</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-submenu index="Traefik">
				<template #title>
					<i class="icon-Partition fix-icon" style="font-size: 18px;"></i>
					<span>Traefik</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="traefikUI">
						<template #title>Traefik UI</template>
					</el-menu-item>
					<el-menu-item index="traefikDoc">
						<template #title>Traefik 文档</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-submenu index="Remco">
				<template #title>
					<i class="icon-template fix-icon" style="font-size: 20px;"></i>
					<span>Remco</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="remcoDoc">
						<template #title>remcoDoc 文档</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-submenu index="Consul">
				<template #title>
					<i class="icon-cluster fix-icon" style="font-size: 18px"></i>
					<span>Consul</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="ConsulManager">
						<template #title>Consul 集群管理</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-menu-item index="Portainer">
				<i class="el-icon-monitor"></i>
				<template #title>Portainer</template>
			</el-menu-item>
			<el-submenu index="Docker">
				<template #title style="text-align: center">
					<i class="icon-docker fix-icon"></i>
					<span>Docker</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="DockerRepositories">
						<template #title>Docker 私有仓库</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-submenu index="MinIO">
				<template #title>
					<i class="icon-group fix-icon"></i>
					<span>MinIO 集群</span>
				</template>
				<el-menu-item-group>
					<el-menu-item index="MinIO_3.2:20001">
						<template #title>3.2:20001 [nginx]</template>
					</el-menu-item>
					<el-menu-item index="MinIO_3.2:20000">
						<template #title>3.2:20000</template>
					</el-menu-item>
					<el-menu-item index="MinIO_3.3:20000">
						<template #title>3.3:20000</template>
					</el-menu-item>
					<el-menu-item index="MinIO_3.4:20000">
						<template #title>3.4:20000</template>
					</el-menu-item>
					<el-menu-item index="MinIO_2.11:20000">
						<template #title>2.11:20000</template>
					</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-menu-item index="vue">
				<i class="icon-bug fix-icon"></i>
				<template #title>Vue3 + Ts 测试页面</template>
			</el-menu-item>
		</el-menu>
	</div>
</template>

<script lang="ts">
	import {computed, defineComponent, ref} from 'vue';
	import {globalStateActions}             from "@/main";

	export default defineComponent({
		emits: [
			"collapse", "update:collapse"
		],
		props: {
			collapse: {
				type: Boolean,
				required: true
			},
		},
		setup(props, context) {
			let collapse = ref<boolean>(props.collapse);
			let isCollapse = computed({
				get: () => {
					return collapse.value;
				},
				set: (v) => {
					context.emit("update:collapse", v);
					collapse.value = v;
				},
			});
			return {
				collapse,
				isCollapse
			};
		},
		methods: {
			handleOpen(key: any, keyPath: any) {
				console.log(key, keyPath);
			},
			handleClose(key: any, keyPath: any) {
				console.log(key, keyPath);
			},
			handleSelect(key: any, keyPath: any) {
				console.log("handleSelect", key, keyPath);
				switch (key) {
					case "home":
						window.location.href = "/";
						globalStateActions.setGlobalState({"navIndex": "/"});
						break;
					case "ncda":
						window.open(window.location.origin + "/ncda", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "ncda"});
						break;
					case "vue":
						window.open(window.location.origin + "/vue", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "vue"});
						break;
					case "traefikUI":
						window.open(window.location.origin + "/traefikUI", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "traefikUI"});
						break;
					case "traefikDoc":
						window.open(window.location.origin + "/traefikDoc", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "traefikDoc"});
						break;
					case "remcoDoc":
						window.open(window.location.origin + "/remcoDoc", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "remcoDoc"});
						break;
					case "ConsulManager":
						window.open(window.location.origin + "/consulManager", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "consulManager"});
						break;
					case "Portainer":
						window.open(window.location.origin + "/portainer", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "Portainer"});
						break;
					case "DockerRepositories":
						window.open(window.location.origin + "/dockerRepositories", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "DockerRepositories"});
						break;
					case "MinIO_3.2:20001":
						window.open(window.location.origin + "/minIO?target=192.168.3.2:20001", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "MinIO"});
						break;
					case "MinIO_3.2:20000":
						window.open(window.location.origin + "/minIO?target=192.168.3.2:20000", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "MinIO"});
						break;
					case "MinIO_3.3:20000":
						window.open(window.location.origin + "/minIO?target=192.168.3.3:20000", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "MinIO"});
						break;
					case "MinIO_3.4:20000":
						window.open(window.location.origin + "/minIO?target=192.168.3.4:20000", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "MinIO"});
						break;
					case "MinIO_2.11:20000":
						window.open(window.location.origin + "/minIO?target=192.168.2.11:20000", "_SELF");
						globalStateActions.setGlobalState({"navIndex": "MinIO"});
						break;
				}
			},
			collapseChange() {
				this.isCollapse = !this.isCollapse;
			},
		}
	});
</script>

<style scoped>
	.fix-icon {
		padding-left: 3px;
		margin-right: 5px;
	}

	.el-menu-vertical {
		height: 100%;
	}

	.el-menu-vertical:not(.el-menu--collapse) {
		width: 200px;
	}

	.nav-bar {
		text-align: left;
		position: fixed;
		height: 100vh;
		z-index: 9999999;
		top: 0;
		left: 0;

		transition: none ease-in-out 1.2s;
	}
</style>
