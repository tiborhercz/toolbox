<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>Convert JSON to YAML</h1>
      </v-col>
    </v-row>
    <json-to-yaml />
  </v-container>
</template>

<script>
import JsonToYaml from '@/components/tools/JsonToYamlComponent'

export default {
  name: 'JsonToYamlView',
  components: {
    JsonToYaml,
  },
}
</script>
