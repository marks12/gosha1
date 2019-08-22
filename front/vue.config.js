module.exports = {
    devServer: {
        headers: {
            "Access-Control-Allow-Origin":"\*"
        },
    },
    assetsDir: process.env.NODE_ENV === 'production'
        ? 'static'
        : 'http://localhost:8080/',
    publicPath: process.env.NODE_ENV === 'production'
        ? '/' // mess up assetsDir if this is blank
        : 'http://localhost:8080/',
};