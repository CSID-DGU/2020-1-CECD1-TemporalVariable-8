<template>
  <RestaurantList></RestaurantList>
</template>

<script lang="ts">
import RestaurantList from "@/components/RestaurantList.vue"

import {Vue, Component} from 'vue-property-decorator'

@Component({
  name: "Search",
  components: {
    RestaurantList
  },
})
export default class Search extends Vue {
}
</script>

<style scoped lang="scss">
</style>