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
        <v-list-item-icon>
          <v-icon>{{ link.icon }}</v-icon>
        </v-list-item-icon>

        <v-list-item-title v-text="link.title" />
      </v-list-item>

      <v-list-group
        v-else
        v-bind:key="link.title"
        no-action
        v-bind:prepend-icon="link.icon"
        v-bind:value="true"
      >
        <template v-slot:activator>
          <v-list-item-title>{{ link.title }}</v-list-item-title>
        </template>

        <v-list-item
          v-for="sublink in link.subLinks"
          v-bind:key="sublink.title"
          v-bind:to="sublink.to"
        >
          <v-list-item-icon>
            <v-icon>{{ sublink.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-title>{{ sublink.title }}</v-list-item-title>
        </v-list-item>
      </v-list-group>
    </div>
  </v-list>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          title: 'Base64',
          icon: 'mdi-code-braces',
          subLinks: [
            {
              title: 'Encode',
              to: '/base64/encode',
            },
            {
              title: 'Decode',
              to: '/base64/decode',
            },
          ],
        },
        {
          title: 'CIDR Calculator',
          icon: 'mdi-web',
          to: '/cidr',
        },
        {
          title: 'JWT Decode',
          icon: 'mdi-code-braces',
          to: '/jwt',
        },
        {
          title: 'JSON Beautify',
          icon: 'mdi-code-braces',
          to: '/json',
        },
      ],
    }
  },
}
</script>
