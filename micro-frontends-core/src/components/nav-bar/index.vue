<template>
	<div class="nav-bar">
		<el-menu default-active="1-4-1" class="el-menu-vertical" @open="handleOpen" @close="handleClose" :collapse="isCollapse">
			<div>
				<el-button style="border: none;width: 100%"
				           :icon="isCollapse?'el-icon-s-unfold':'el-icon-s-fold'"
				           @click="collapseChange">
							{{isCollapse?'':'收起导航栏'}}
				</el-button>
			</div>

			<el-submenu index="1">
				<template #title>
					<i class="el-icon-location"></i>
					<span>导航一</span>
				</template>
				<el-menu-item-group>
					<template #title>分组一</template>
					<el-menu-item index="1-1">选项1</el-menu-item>
					<el-menu-item index="1-2">选项2</el-menu-item>
				</el-menu-item-group>
				<el-menu-item-group title="分组2">
					<el-menu-item index="1-3">选项3</el-menu-item>
				</el-menu-item-group>
				<el-submenu index="1-4">
					<template #title>选项4</template>
					<el-menu-item index="1-4-1">选项1</el-menu-item>
				</el-submenu>
			</el-submenu>
			<el-menu-item index="2">
				<i class="el-icon-menu"></i>
				<template #title>导航二</template>
			</el-menu-item>
			<el-menu-item index="3" disabled>
				<i class="el-icon-document"></i>
				<template #title>导航三</template>
			</el-menu-item>
			<el-menu-item index="4">
				<i class="el-icon-setting"></i>
				<template #title>导航四</template>
			</el-menu-item>
		</el-menu>
	</div>
</template>

<script>
	export default {
		name: "navBar",
		props:{
			collapse:{
				type:Boolean,
				default: true
			}
		},
		computed:{
			isCollapse:{
				get:function(){
					return this.collapse;
				},
				set:function(v){
					this.$emit("update:collapse",v);
				}
			}
		},
		data() {
			return {};
		},
		methods: {
			handleOpen(key, keyPath) {
				console.log(key, keyPath);
			},
			handleClose(key, keyPath) {
				console.log(key, keyPath);
			},
			collapseChange() {
				this.isCollapse = !this.isCollapse;
			}
		}
	}
</script>

<style scoped>
	.el-menu-vertical{
		height: 100%;
	}
	.el-menu-vertical:not(.el-menu--collapse) {
		width: 200px;
	}

	.nav-bar{
		position: fixed;
		height: 100vh;
		z-index: 9999999;
		top: 0;
		left: 0;
	}
</style>
