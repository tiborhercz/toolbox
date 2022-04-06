<template>
  <div>
<!--eslint-disable-->
<pre v-bind:ref="'prism'">
<code
  class="language-json"
  v-text="content"
/>
</pre>
  <!--eslint-enable-->
  </div>
</template>

<script>
import Prism from 'prismjs'
import 'prismjs/components/prism-json'

export default {
  name: 'JsonFormatter',
  components: {},
  props: {
    json: {
      default: '',
      type: String,
    },
  },
  computed: {
    content() {
      let jsonValue = this.json

      this.$nextTick(() => {
        Prism.highlightAllUnder(this.$refs.prism)
      })

      try {
        jsonValue = JSON.stringify(JSON.parse(jsonValue), null, 4)

        this.$emit('formattedJson', jsonValue)
        return jsonValue
      } catch (e) {
        if (this.json !== '') {
          this.$emit('errorMessage', e.message)
        }

        return JSON.stringify({})
      }
    },
  },
}
</script>

<style scoped>
.v-application code {
  padding: 0;
}
</style>
