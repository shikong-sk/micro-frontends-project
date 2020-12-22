<template>
	<div ref="container">
		<div id="nav">
			<router-link to="/">Home</router-link>
			|
			<router-link to="/about">About</router-link>
			|
		</div>
		<div class="home">
			<img alt="Vue logo" src="../assets/logo.png">
			<HelloWorld msg="Welcome to Your Vue.js + TypeScript App"/>
		</div>

		<el-row>
			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>Traefik</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/traefikUI">Traefik UI</el-link>
						</div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/traefikDoc">Traefik 文档</el-link>
						</div>
					</div>
				</el-card>
			</el-col>

			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>RemCo</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/remCoDoc">RemCo 文档</el-link>
						</div>
					</div>
				</el-card>
			</el-col>

			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>Consul</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/consulManager">Consul 管理</el-link>
						</div>
					</div>
				</el-card>
			</el-col>

			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>Portainer</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/portainer">Portainer 面板</el-link>
						</div>
					</div>
				</el-card>
			</el-col>

			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>Docker</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/dockerRepositories">Docker 私有仓库</el-link>
						</div>
					</div>
				</el-card>
			</el-col>

			<el-col :lg="4">
				<el-card shadow="hover">
					<template #header>
						<div>
							<span>MinIO</span>
						</div>
					</template>
					<div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/minIO?target=192.168.3.2:20001">
								192.168.3.2:20001 nginx
							</el-link>
						</div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/minIO?target=192.168.3.2:20000">
								192.168.3.2:20000
							</el-link>
						</div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/minIO?target=192.168.3.3:20000">
								192.168.3.3:20000
							</el-link>
						</div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/minIO?target=192.168.3.4:20000">
								192.168.3.4:20000
							</el-link>
						</div>
						<div class="link-block">
							<el-link type="primary" target="_self" href="/minIO?target=192.168.2.11:20000">
								192.168.2.11:20000
							</el-link>
						</div>
					</div>
				</el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script lang="ts">
    import {defineComponent, getCurrentInstance, onMounted, ref} from "vue";
    import useNavBar                                             from "@/components/nav-bar/UseNavBar";
    import {globalStateActions}                                  from "@/main";

    export default defineComponent({
        name: "TraefikDoc",
        setup: (prop, context) => {
            let container = ref();
            console.log(container);
            useNavBar(container);


            // 获取当前内部组件实例
            let instance = getCurrentInstance();
            // 获取全局属性/方法
            let ctx = instance?.appContext.config.globalProperties;

            let getNetWork = () => {
                ctx?.$Api.systemApi.getNetWork().then((response: any) => {
                    let res = (response.data as InterfaceInfo[]);
                    console.log(res);
                    for (let i = 0; i < res.length; i++) {
	                    let interfaceInfo = res[i];
                    }
                });
            };

            type InterfaceInfo = {
                name: string,
                mac: string,
                flags: string,
                ipAddrs: InterfaceAddrInfo[]
            }
            type InterfaceAddrInfo = {
                ip: string,
                mask: string,
                maskHex: string
            }

            onMounted(() => {
                globalStateActions.setGlobalState({"navIndex": "home"});
                getNetWork();
            });
            return {
                container,
                ctx,
                getNetWork
            };
        }
    });
</script>

<style scoped>
	.link-block {
		height: 25px;
	}
</style>
