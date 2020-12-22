module.exports = {
    devServer: {
        // 在基座配置 端口要和这里一样。
        port: 10000,
        //  任何人都能访问的服务器
        headers: {
            "Access-Control-Allow-Origin": "*",
        },
        /**
         * 如果在dev模式下运行主框架 若子项目中有使用 代理进行跨域访问的 那么主框架下也需要配置相应内容
         * 或者 搭建一个 nginx 服务器 专门转发请求
         */
        proxy: {
            // pathRewrite 中将 /ncda/api/ 替换为空
            "/ncda/api/": {
                //配置代理地址，前端请求的所有接口都需要带的前缀
                target:"http://localhost:18848/ncda/",
                // target: "http://192.168.1.241:1080/ncda/",
                // target: 'http://192.168.2.5:8010/ncda/',
                changeOrigin: true, //是否进行跨域
                secure: false,
                pathRewrite: {
                    //我使用了 nginx 作为反向代理，所以，需要把前缀替换为nginx 配置中的代理路径
                    "^/ncda/api/": "", //服务器请求地址中，若没有/api ，则替换为 /
                },
            },
            "/ncdk/api/": {
                target: "http://192.168.1.241:1089/ncdk/",
                // target: "http://192.168.2.3:7010/ncdk/",
                changeOrigin: true, //是否进行跨域
                secure: false,
                pathRewrite: {
                    "^/ncdk/api/": "",
                },
            },
            "/relation/api": {
                target: "http://192.168.0.103:8088/analysis/",
                changeOrigin: true, //是否进行跨域
                secure: false,
                pathRewrite: {
                    "/relation/api/": "",
                },
            },
            "/hotspot/api/": {
                target: "http://192.168.0.106:8018/hotspot/",
                changeOrigin: true, //是否进行跨域
                secure: false,
                pathRewrite: {
                    "/hotspot/api/": "",
                },
            },
            "/statistics/api/": {
                target: "http://192.168.0.106:8028/statistics/",
                changeOrigin: true, //是否进行跨域
                secure: false,
                pathRewrite: {
                    "/statistics/api/": "",
                },
            },
            "^/config/": {
                target: "http://192.168.1.241:8999",
            },
            "^/base/": {
                target: "http://localhost:18848/base/",
                changeOrigin: true, //是否进行跨域
                pathRewrite: {
                    "^/base/": "",
                },
            },
            "^/system/": {
                target: "http://localhost:18848/system/",
                changeOrigin: true, //是否进行跨域
                pathRewrite: {
                    "^/system/": "",
                },
            },
        }
    }
}
