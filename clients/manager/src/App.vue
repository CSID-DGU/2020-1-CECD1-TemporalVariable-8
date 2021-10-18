<template>
  <v-app>
    <v-app-bar
        app
        color="primary"
        dark
    >
      <div class="d-flex align-center">
        <router-link to="/">
          <v-img
              alt="Vuetify Logo"
              class="shrink mr-2"
              contain
              :src="require('@/assets/Logo.png')"
              transition="scale-transition"
              width="40"
          />
        </router-link>

        <v-card-title>맛잇고</v-card-title>
      </div>

      <v-spacer></v-spacer>

      <v-btn
          href="/auth/logout"
          target="_blank"
          text
          v-if="isLogin"
      >
        <span class="mr-2">로그아웃</span>
        <v-icon>mdi-logout</v-icon>
      </v-btn>
      <v-btn
          href="/auth/login/google"
          target="_blank"
          text
          v-else
      >
        <span class="mr-2">로그인</span>
        <v-icon>mdi-login</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <v-card-title v-if="!isLogin">음식점을 보려면 로그인해주세요!</v-card-title>
      <router-view v-else ref="/"/>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: 'App',

  mounted() {
    this.loadSession()
  },
  methods: {
    loadSession: function () {
      return this.$axios.get('/auth/self')
          .then(res => {
            this.session = res.data
          })
          .catch(() => {
            this.session = undefined
          })

      // 디버깅시 로그인
      // let dat = {data: {client: {id: 1}}}
      // return Promise.resolve(dat)
      //     .then(res => {
      //       this.session = res.data
      //     })
      //     .catch(() => {
      //       this.session = undefined
      //     })
    }
  },
  computed: {
    isLogin: function () {
      try {
        return this.session.client.id !== undefined;
      } catch (e) {
        return false
      }
    }
  },
  data: () => ({
    session: null
  }),
};
</script>
