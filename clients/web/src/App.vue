<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        />

        <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
          width="100"
        />
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-layout fill-height>
        <v-navigation-drawer permanent>
          <v-toolbar flat class="transparent">
            <v-list class="pa-0">
              <v-list-item avatar>
                <v-list-item-avatar>
                  <img src="https://randomuser.me/api/portraits/men/85.jpg">
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title>John Leider</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-toolbar>
          <v-divider/>
          <v-list>
            <v-list-item to="/">
              <v-list-item-content>Home</v-list-item-content>
            </v-list-item>

            <v-list-item to="/about">
              <v-list-item-content>About</v-list-item-content>
            </v-list-item>
          </v-list>
        </v-navigation-drawer>
        <router-view/>
      </v-layout>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'App',

  components: {
  },

  data: () => ({
  }),
});
</script>
