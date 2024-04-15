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
    indent: {
      default: 2,
      type: Number,
    },
    value: {
      default: '',
      type: String,
    },
  },
  computed: {
    content() {
      let jsonValue = this.value

      this.$nextTick(() => {
        Prism.highlightAllUnder(this.$refs.prism)
      })

      try {
        jsonValue = JSON.stringify(JSON.parse(jsonValue), null, this.getIndentValue())

        this.$emit('formattedValue', jsonValue)
        return jsonValue
      } catch (e) {
        if (this.value !== '') {
          this.$emit('errorMessage', e.message)
        }

        return JSON.stringify({})
      }
    },
  },
  methods: {
    getIndentValue() {
      const indentValue = this.indent
      if (!indentValue) {
        return 2
      }

      return this.indent
    },
  },
}
</script>

<style scoped>
.v-application code {
  padding: 0;
}
</style>
