module.exports = {
	devServer:{
		// 在基座配置 端口要和这里一样。
		port:10000,
		//  任何人都能访问为的服务器
		headers:{
			"Access-Control-Allow-Origin":"*"
		},
		/**
		 * 如果在dev模式下运行主框架 若子项目中有使用 代理进行跨域访问的 那么主框架下也需要配置相应内容
		 * 或者 搭建一个 nginx 服务器 专门转发请求
		 */
		proxy: {
			"/ncda/api/": {
				target: "http://localhost:8848/ncda/",
				// target: "http://192.168.1.241:1080/ncda/",
				// target: 'http://192.168.2.5:8010/ncda/',
				changeOrigin: true, //是否进行跨域
				secure: false,
				pathRewrite: {
					"^/ncda/api/": "",
				},
			},
		}
	},
}
