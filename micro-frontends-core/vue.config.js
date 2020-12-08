module.exports = {
	devServer:{
		// 在基座配置 端口要和这里一样。
		port:10000,
		//  任何人都能访问为的服务器
		headers:{
			"Access-Control-Allow-Origin":"*"
		}
	},
}