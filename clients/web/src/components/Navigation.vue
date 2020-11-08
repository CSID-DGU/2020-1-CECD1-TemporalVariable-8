<template>
  <v-app-bar
      app
      dark
      src="https://picsum.photos/1920/1080?random"
      fade-img-on-scroll
      shrink-on-scroll
      color="primary"
      height="340"
      extended
      extension-height="52"
  >
    <!-- 네비게이션 뒤의 이미지에 그라디언트를 넣는 부분, 딱히 그 외엔 용도 없음 -->
    <template v-slot:img="{ props }">
      <v-img v-bind="props" gradient="to top right, rgba(55,236,186,.7), rgba(25,32,72,.7)"/>
    </template>
    <!-- 네비게이션바의 주 파트 -->
    <v-row class="mx-auto fill-height" style="max-width: 1920px">
      <v-col
          class="pa-0 fill-height"
          md="2"
          sm="3"
          cols="6"
      >
        <div class="d-flex fill-height justify-start">
          <v-btn icon>
            <v-icon>mdi-menu</v-icon>
          </v-btn>
        </div>
      </v-col>
      <v-col
          class="pa-0 fill-height nav-searcher"
          md="8"
          sm="6"
          cols="0"
      >
        <div class="d-flex fill-height">
          <v-text-field
              class="my-auto pt-1"
              label="위치 알 수 없음"
              prepend-inner-icon="mdi-map-marker"
              outlined
              filled
              dense
          />
        </div>
      </v-col>
      <v-col
          class="pa-0 fill-height"
          md="2"
          sm="3"
          cols="6"
      >
        <div class="d-flex fill-height justify-end">
          <v-avatar v-if="isLogin">
            <v-img
                :src="require('@/assets/logo.png')"
                style="cursor: pointer;border: 1px solid lightgrey"
                @click="say"
            />
          </v-avatar>
          <LoginDialog v-else>
            <template v-slot:activator="{on, attrs}">
              <v-btn
                  icon
                  v-bind="attrs"
                  v-on="on"
              >
                <v-icon>mdi-account-key-outline</v-icon>
              </v-btn>
            </template>
          </LoginDialog>
          <!-- 로그인 창 -->
        </div>
      </v-col>
    </v-row>

    <!-- 네비게이션 추가 파트의 탭 버튼 -->
    <template v-slot:extension>
      <v-container>
        <v-row
            class="mx-auto"
            style="max-width: 1920px"
        >
          <v-tabs
              background-color="transparent"
              grow
              centered
              icons-and-text
              height="52"
          >
            <v-tabs-slider/>
            <v-tab>
              {{ $t('nav.search') }}
              <v-icon>mdi-magnify</v-icon>
            </v-tab>
            <v-tab>
              {{ $t('nav.favorite') }}
              <v-icon>mdi-star</v-icon>
            </v-tab>
            <v-tab>
              {{ $t('nav.history') }}
              <v-icon>mdi-history</v-icon>
            </v-tab>
          </v-tabs>
        </v-row>
      </v-container>
    </template>
  </v-app-bar>
</template>
<style scoped lang="scss">
.nav-searcher {
  @media (max-width: 600px) {
    display: none;
  }
}
</style>
<script lang="ts">
import {Vue, Component} from "vue-property-decorator"
import {mapState, mapGetters} from "vuex"
import LoginDialog from "./dialogs/Login.vue"

@Component({
  name: "Navigation",
  components: {
    LoginDialog
  },
  computed: {
    ...mapState("login", ['name']),
    ...mapGetters("login", ['isLogin']),
  },

})
export default class Navigation extends Vue {
  loginDialog = false

  say() {
    alert('hello')
  }
}
</script>
