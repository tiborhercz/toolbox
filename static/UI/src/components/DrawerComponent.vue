<script>
export default {
  data() {
    return {
      items: [
        {
          title: 'Base64 Encode',
          icon: 'mdi-code-braces',
          to: '/base64/encode',
        },
        {
          title: 'Base64 Decode',
          icon: 'mdi-code-braces',
          to: '/base64/decode',
        },
        {
          title: 'CIDR Calculator',
          icon: 'mdi-web',
          to: '/cidr',
        },
        {
          title: 'JWT Decode',
          icon: 'mdi-code-braces',
          to: '/json/jwt',
        },
        {
          title: 'JSON Beautify',
          icon: 'mdi-code-braces',
          to: '/json/json-beautify',
        },
        {
          title: 'JSON to YAML',
          icon: 'mdi-code-braces',
          to: '/json/json-to-yaml',
        },
        {
          title: 'Hash',
          icon: 'mdi-pound',
          to: '/hash',
        },
      ],
    }
  },
}
</script>

<template>
  <v-list
    dense
  >
    <div
      v-for="(link, i) in items"
      v-bind:key="i"
    >
      <v-list-item
        v-if="!link.subLinks"
        v-bind:to="link.to"
        class="v-list-item"
      >
        <template v-slot:prepend>
          <v-icon
            v-bind:icon="link.icon"
          />
        </template>

        <v-list-item-title
          v-text="link.title"
        />
      </v-list-item>

      <v-list-group
        v-else
        v-bind:key="link.title"
        no-action
        v-bind:prepend-icon="link.icon"
        v-bind:value="true"
      >
        <template v-slot:activator="{ props }">
          <v-list-item
            v-bind="props"
            v-bind:title="link.title"
          />
        </template>

        <v-list-item
          v-for="sublink in link.subLinks"
          v-bind:key="sublink.title"
          v-bind:to="sublink.to"
        >
          <v-icon>{{ sublink.icon }}</v-icon>
          <v-list-item-title>{{ sublink.title }}</v-list-item-title>
        </v-list-item>
      </v-list-group>
    </div>
  </v-list>
</template>
