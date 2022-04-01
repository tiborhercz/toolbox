<template>
  <v-container>
    <v-row>
      <v-col>
        <h1>CIDR calculator</h1>
      </v-col>
    </v-row>
    <Cidr />
  </v-container>
</template>

<script>
import Cidr from '@/components/tools/Cidr'

export default {
  name: 'CidrView',
  components: {
    Cidr,
  },
  mounted() {},
}
</script>
