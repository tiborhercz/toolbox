import { createApp } from 'vue'
import router from './router'

import { registerPlugins } from './plugins'

import App from './App.vue'
import Go from './wasm_exec'

const go = new Go()

const app = createApp(App)

registerPlugins(app)

app.use(router)

app.mount('#app')

WebAssembly.instantiateStreaming(fetch('/main.wasm'), go.importObject).then((result) => {
  go.run(result.instance)
})
