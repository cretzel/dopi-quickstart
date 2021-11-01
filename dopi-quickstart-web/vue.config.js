const devServer = require('./src/mock/devServer')

module.exports = {
    publicPath: "/quickstart",
    devServer: {
        before: devServer,
        proxy: {
            '/api/quickstart': {
                target: 'http://localhost:8081'
            }
        }
    }
}