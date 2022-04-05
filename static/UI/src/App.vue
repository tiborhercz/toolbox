<template>
  <v-app v-bind:style="{background: $vuetify.theme.themes[darkTheme].background}">
    <v-navigation-drawer
      v-bind:color="$vuetify.theme.themes[darkTheme].drawer"
      app
    >
      <v-container>
        <v-row>
          <v-col>
            <h1>Toolbox</h1>
          </v-col>
        </v-row>
      </v-container>
      <Drawer />
      <template v-slot:append>
        <div class="pa-2">
          <v-switch
            v-model="darkMode"
            label="Dark theme"
            v-on:change="toggleDarkMode"
          />
        </div>
      </template>
    </v-navigation-drawer>
    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Drawer from '@/components/Drawer'

export default {
  name: 'App',
  components: { Drawer },
  data: () => ({
    darkMode: true,
  }),
  computed: {
    darkTheme() {
      return this.darkMode ? 'dark' : 'light'
    },
  },
  mounted() {
    const theme = localStorage.getItem('dark_theme')
    if (theme) {
      this.$vuetify.theme.dark = theme === 'true'
      this.darkMode = theme === 'true'
    }
  },
  methods: {
    toggleDarkMode() {
      if (this.darkMode) {
        window.document.body.style.backgroundColor = '#222222'
      } else {
        window.document.body.style.backgroundColor = '#ffffff'
      }
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark
      localStorage.setItem('dark_theme', this.$vuetify.theme.dark.toString())
    },
  },
}
</script>
