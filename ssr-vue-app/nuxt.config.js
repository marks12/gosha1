module.exports = {
  server: {
    port: 5000, // default: 3000
    host: '0.0.0.0', // default: localhost
  },
  /*
  ** Headers of the page
  */
  head: {
    title: 'Shop application',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Shop application description' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** Run ESLint on save
    */
    // extend (config, { isDev, isClient }) {
    //   if (isDev && isClient) {
    //     config.module.rules.push({
    //       enforce: 'pre',
    //       test: /\.(js|vue)$/,
    //       loader: 'eslint-loader',
    //       exclude: /(node_modules)/
    //     })
    //   }
    // },
    vendor: ['axios', 'vue-awesome-swiper'],
  },
  css: [
    '~assets/scss/style.scss'
  ],
  modules: ['@nuxtjs/style-resources'],
  styleResources: {
    scss: [
      'assets/scss/_variables.scss',
    ]
  },
  plugins: [
    { src: '~/plugins/ws.js', mode: 'client' },
    { src: '~/plugins/router-interceptor.js', mode: 'client' },
    { src: '~/plugins/vue-moment.js' }
  ],
  router: {
    extendRoutes (routes, resolve) {
      routes.push({
        name: 'custom',
        path: '*',
        component: resolve(__dirname, 'pages/index.vue')
      })
    }
  }
};
