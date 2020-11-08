<template>
  <v-app>
    <Navigation/>
    <v-main>
      <v-container fluid class="main-router-view">
        <v-row>
          <v-col class="pa-0">
            <router-view ref="/search"/>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>
<style lang="scss">
.main-router-view{
  max-width: 1920px;
}
</style>
<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import HelloWorld from './components/HelloWorld.vue';
import Navigation from './components/Navigation.vue';

@Component({
  name: 'App',
  components: {
    HelloWorld,
    Navigation,
  }
})
export default class App extends Vue {
  // data 속성
  message = 'hello';

  // methods 속성
  sayHi() {
    console.log(this.message);
  }
}
// export default Vue.extend({
//   name: 'App',
//
//   components: {
//     HelloWorld,
//   },
//
//   data: () => ({
//     //
//   }),
// });
</script>
