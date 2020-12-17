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
			<el-submenu index="1">
				<template #title>
					<i class="el-icon-location"></i>
					<span>微前端1</span>
				</template>
				<el-menu-item-group>
					<template #title>微前端1</template>
					<el-menu-item index="ncda">系统1</el-menu-item>
				</el-menu-item-group>
			</el-submenu>
			<el-menu-item index="vue">
				<i class="el-icon-menu"></i>
				<template #title>导航二</template>
			</el-menu-item>
		</el-menu>
	</div>
</template>

<script lang="ts">
	import {computed, defineComponent, ref} from 'vue';
	import {globalStateActions}             from "@/main";

	export default defineComponent({
		emits:[
			"collapse","update:collapse"
		],
		props:{
			collapse: {
				type:Boolean,
				required: true
			},
		},
		setup(props,context){
			let collapse = ref<boolean>(props.collapse)
			let isCollapse = computed({
				get:  () => {
					console.log("collapse computed get",collapse.value);
					return collapse.value;
				},
				set:  (v)=> {
					console.log("collapse computed", v);
					context.emit("update:collapse", v);
					collapse.value = v;
				},
			})
			return {
				collapse,
				isCollapse
			}
		},
		methods:{
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
						// this.$router.push({name:"Home"});
						// window.open(window.location.origin + "/","_SELF");
						window.location.href = "/";
						break;
					case "ncda":
						window.open(window.location.origin + "/ncda","_SELF");
                        this.$router.push({path:"/ncda"})
						globalStateActions.setGlobalState({"navIndex":"ncda"});
						break;
					case "vue":
						window.open(window.location.origin + "/vue","_SELF");
                        globalStateActions.setGlobalState({"navIndex":"ncda"});
						// this.$router.push({path:"/vue"})
						break;
				}
			},
			collapseChange() {
				console.log("this.isCollapse.value ",this.isCollapse );
				this.isCollapse = !this.isCollapse;
				console.log("this.isCollapse.value ",this.isCollapse );
			},
		}
	})
</script>

<style scoped>
	.el-menu-vertical {
		height: 100%;
	}

	.el-menu-vertical:not(.el-menu--collapse) {
		width: 200px;
	}

	.nav-bar {
		position: fixed;
		height: 100vh;
		z-index: 9999999;
		top: 0;
		left: 0;
	}
</style>
