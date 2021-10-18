<template>
  <v-container>
    <v-row>
      <v-col cols="4" v-for="rest in restaurants" :key="rest.id">
        <Restaurant :rest_id="rest.id" :rest_name="rest.name" :rest_description="rest.desc"
                    :rest_avatar="rest.avatar"/>
      </v-col>
      <v-col cols="4">
        <v-card
            class="mx-auto my-12"
            height="240"
            elevation="2"
        >
          <router-link to="/add">
            <v-container fill-height fluid>
              <v-row align="center" justify="center">
                <v-img
                    class="me-auto"
                    height="120"
                    contain
                    :src="require('@/assets/add.svg')"
                />
              </v-row>
              <v-row align="center" justify="center">
                추가하기
              </v-row>
            </v-container>
          </router-link>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// @ is an alias to /src
import Restaurant from '@/components/Restaurant.vue'

export default {
  name: 'Home',
  components: {
    Restaurant
  },
  mounted() {
    this.loadRestaurants()
  },
  methods: {
    loadRestaurants() {
      this.$axios.get("/api/restaurants/")
          .then((res) => {
            for (let p of res.data.names) {
              if (p.url === undefined || p.url.length === 0) {
                this.$axios.get("/api/restaurants/" + p.path.join("/") + "/")
                    .then((res) => {
                      this.restaurants.push({
                        id: res.data.id,
                        name: res.data.name.path,
                        desc: res.data.description,
                        avatar: res.data.avatar,
                      })
                    })
                    .catch(function (err) {
                      console.log(err)
                    })
              }
            }
          })
          .catch(function (err) {
            console.log(err)
          })
      // this.restaurants = [
      //   {
      //     id : "123",
      //     name : ["동국반점"],
      //     desc : "맛있는 짜장면 있어요!",
      //     avatar : "jjajjangmyoun1.jpg",
      //   },
      //   {
      //     id : "124",
      //     name : ["동국반점",  "충무로점"],
      //     desc : "동국반점 충무로점입니다. 맛있는 짬뽕도 있어요!",
      //     avatar : "jjajjangmyoun2.jpg",
      //   }
      // ]
    }
  },
  data: () => ({
    restaurants: []
  })
}
</script>
