import Vue from 'vue'

import App from './App'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'

import Go from './wasm_exec'

Vue.config.productionTip = false

const go = new Go()
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
  go.run(result.instance)

  new Vue({
    router,
    store,
    vuetify,
    render: (h) => h(App),
  }).$mount('#app')
})
