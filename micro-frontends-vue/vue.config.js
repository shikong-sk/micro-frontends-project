module.exports = {
	devServer: {
		port: 10001,
		// 允许跨域
		headers: {
			"Access-Control-Allow-Origin": "*",
		},
	},
	//  配置打包后的模块， 因为基座引用的是打包后的
	configureWebpack: {
		// 将应用打包为 umd 库格式
		output: {
			library: `micro-frontends-vue`,
			libraryTarget: "umd",
			jsonpFunction:`webpackJsonp_micro-frontends-vue`
		},
	},
};
