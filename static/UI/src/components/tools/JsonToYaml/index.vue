<template>
  <v-row>
    <v-col
      cols="12"
    >
      <p>
        JSON to YAML
      </p>
      <v-row>
        <v-col
          cols="10"
          md="7"
        >
          <v-textarea
            ref="textarea"
            v-model="inputValue"
            name="input-7-1"
            label="Input"
            v-bind:error="error"
            v-bind:error-messages="errorMessages"
            outlined
            v-on:keydown.tab.exact.prevent
            v-on:keydown.tab.exact="insertTabChar"
          />
        </v-col>
        <v-col
          cols="2"
          md="2"
        >
          <v-text-field
            v-model.number="indent"
            label="Indent"
            outlined
          />
        </v-col>
      </v-row>
      <yaml-formatter
        v-if="yamlValue"
        v-bind:value="yamlValue"
        v-bind:indent="indent"
        v-on:formattedValue="setFormattedYaml"
        v-on:errorMessage="setErrorMessage"
      />
    </v-col>
  </v-row>
</template>

<script>
import JsonToYaml from 'json-to-pretty-yaml'
import yamlFormatter from '@/components/CodeFormatters/YamlFormatter/index'

export default {
  name: 'JsonToYaml',
  components: {
    yamlFormatter,
  },
  props: {},
  data() {
    return {
      value: '',
      indent: 2,
      jsonValue: '',
      yamlValue: '',
      formattedValue: {},
      error: false,
      errorMessages: [],
    }
  },
  computed: {
    inputValue: {
      get() {
        return this.value
      },
      async set(newValue) {
        if (this.isJson(newValue)) {
          this.yamlValue = JsonToYaml.stringify(JSON.parse(newValue))
          this.jsonValue = ''
        }

        this.value = newValue
        this.error = false
        this.errorMessages = []
      },
    },
  },
  methods: {
    isJson(str) {
      try {
        JSON.parse(str)
      } catch (e) {
        return false
      }
      return true
    },
    insertTabChar() {
      const ref = this.$refs.textarea
      const startPos = ref.$el.querySelector('textarea').selectionStart
      const endPos = ref.$el.querySelector('textarea').selectionEnd

      const substrStart = this.inputValue.substring(0, startPos)
      const substrEnd = this.inputValue.substring(endPos)

      this.inputValue = `${substrStart}\t${substrEnd}`

      this.$nextTick(() => this.setCursorPosition(ref.$el.querySelector('textarea'), startPos + 1))
    },
    setCursorPosition(el, pos) {
      el.focus()
      el.setSelectionRange(pos, pos)
    },
    setFormattedYaml(value) {
      this.formattedValue = value
    },
    setErrorMessage(errorMessage) {
      this.error = true
      this.errorMessages.push(errorMessage)
    },
  },
}
</script>
