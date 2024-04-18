<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>Hash {{ type }}</h1>
      </v-col>
    </v-row>
    <Hash
      v-bind:type="type"
    />
  </v-container>
</template>

<script>
import Hash from '@/components/tools/HashComponent'

export default {
  name: 'HashView',
  components: {
    Hash,
  },
  props: {
    type: {
      default: '',
      type: String,
    },
  },
}
</script>
