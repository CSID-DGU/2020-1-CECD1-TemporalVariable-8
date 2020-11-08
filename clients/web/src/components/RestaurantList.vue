<template>
  <v-container fluid>
    <v-row dense>
      <v-col cols="3">
        <RestaurantCard/>
      </v-col>
      <v-col cols="6">
        <RestaurantCard/>
      </v-col>
      <v-col cols="3">
        <RestaurantCard/>
      </v-col>
      <v-col cols="3">
        <RestaurantCard/>
      </v-col>
      <v-col cols="3">
        <RestaurantCard/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import RestaurantCard from "@/components/RestaurantCard.vue"
import {Vue, Component} from 'vue-property-decorator'

@Component({
  name: "RestaurantList",
  components: {
    RestaurantCard
  },
})
export default class RestaurantList extends Vue {

}
</script>

<style scoped>

</style>