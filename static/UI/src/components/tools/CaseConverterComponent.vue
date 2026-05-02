<template>
  <v-row>
    <v-col cols="12">
      <p class="mb-4">
        Convert text between different case styles. Type in any field to convert from that case.
      </p>
      <v-row>
        <v-col cols="12" md="8">
          <v-textarea
            v-model="input"
            label="Input"
            variant="outlined"
            rows="3"
            placeholder="Type or paste text here..."
            autofocus
          />
        </v-col>
      </v-row>
      <v-row>
        <v-col
          v-for="c in cases"
          :key="c.key"
          cols="12"
          md="6"
        >
          <v-text-field
            :model-value="c.value"
            :label="c.label"
            variant="outlined"
            density="compact"
            readonly
            append-inner-icon="mdi-content-copy"
            @click:append-inner="copy(c.value)"
          />
        </v-col>
      </v-row>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useSnackbar } from '@/helpers/useSnackbar'

const { showSnackbar } = useSnackbar()

const input = ref('')

function splitWords(str) {
  return str
    .replace(/([a-z\d])([A-Z])/g, '$1 $2')
    .replace(/([A-Z]+)([A-Z][a-z])/g, '$1 $2')
    .replace(/[\s\-_./\\|:]+/g, ' ')
    .trim()
    .split(' ')
    .filter(w => w.length > 0)
    .map(w => w.toLowerCase())
}

function capitalize(w) {
  return w.charAt(0).toUpperCase() + w.slice(1)
}

const cases = computed(() => {
  const words = splitWords(input.value)
  return [
    { key: 'lower',    label: 'Lowercase',    value: input.value.toLowerCase() },
    { key: 'upper',    label: 'Uppercase',    value: input.value.toUpperCase() },
    { key: 'camel',    label: 'Camelcase',    value: words.length ? words[0] + words.slice(1).map(capitalize).join('') : '' },
    { key: 'capital',  label: 'Capitalcase',  value: words.map(capitalize).join(' ') },
    { key: 'constant', label: 'Constantcase', value: words.join('_').toUpperCase() },
    { key: 'dot',      label: 'Dotcase',      value: words.join('.') },
    { key: 'header',   label: 'Headercase',   value: words.map(capitalize).join('-') },
    { key: 'no',       label: 'Nocase',       value: words.join(' ') },
    { key: 'param',    label: 'Paramcase',    value: words.join('-') },
    { key: 'pascal',   label: 'Pascalcase',   value: words.map(capitalize).join('') },
    { key: 'path',     label: 'Pathcase',     value: words.join('/') },
    { key: 'sentence', label: 'Sentencecase', value: words.length ? capitalize(words[0]) + (words.length > 1 ? ' ' + words.slice(1).join(' ') : '') : '' },
    { key: 'snake',    label: 'Snakecase',    value: words.join('_') },
  ]
})

function copy(value) {
  navigator.clipboard.writeText(value)
  showSnackbar('Copied!')
}
</script>
