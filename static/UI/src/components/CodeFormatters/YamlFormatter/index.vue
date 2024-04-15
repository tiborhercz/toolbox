<template>
  <div>
<!--eslint-disable-->
<pre v-bind:ref="'prism'">
<code
  class="language-yaml"
  v-text="content"
/>
</pre>
  <!--eslint-enable-->
  </div>
</template>

<script>
import Prism from 'prismjs'
import 'prismjs/components/prism-yaml'

export default {
  name: 'YamlFormatter',
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
      this.$nextTick(() => {
        Prism.highlightAllUnder(this.$refs.prism)
      })

      try {
        this.$emit('formattedValue', this.value)
        return this.value
      } catch (e) {
        if (this.value !== '') {
          this.$emit('errorMessage', e.message)
        }

        return ''
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
