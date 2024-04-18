<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>JSON beautify</h1>
      </v-col>
    </v-row>
    <json />
  </v-container>
</template>

<script>
import Json from '@/components/tools/JsonBeautifyComponent'

export default {
  name: 'JsonBeautify',
  components: {
    Json,
  },
}
</script>
